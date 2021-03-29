package docsgen

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/stoewer/go-strcase"
)

var (
	helmDocsDir = "content/reference/helm"

	ossFileMapping = map[string]string{
		"codegen/helm/gloo_mesh_helm_values_reference.md":  "%s/%s/gloo_mesh.md",
		"codegen/helm/cert_agent_helm_values_reference.md": "%s/%s/cert_agent.md",
	}

	enterpriseFileMapping = map[string]string{
		"enterprise-networking/codegen/helm/enterprise_networking_helm_values_reference.md": "%s/%s/enterprise_networking.md",
		"enterprise-networking/codegen/helm/enterprise_agent_helm_values_reference.md":      "%s/%s/enterprise_agent.md",
	}

	rbacWwebookFileMapping = map[string]string{
		"rbac-webhook/codegen/chart/rbac_webhook_helm_values_reference.md": "%s/%s/rbac_webhook.md",
	}

	helmValuesIndex = `
---
title: "%s"
description: Reference for Helm values. 
weight: 2
---
{{%% children description="true" %%}}
`
)

func copyHelmValuesDocsForAllCharts(client *github.Client, rootDir string) error {
	// flush root directory for idempotence
	helmDocsDir := filepath.Join(rootDir, helmDocsDir)
	os.RemoveAll(helmDocsDir)
	os.MkdirAll(helmDocsDir, 0755)

	// create root index
	if err := createFileIfNotExists(helmDocsDir+"/"+"_index.md", fmt.Sprintf(helmValuesIndex, "Helm Values Reference")); err != nil {
		return eris.Errorf("error creating Helm values index file: %v", err)
	}

	// Gloo Mesh OSS
	if err := copyHelmValuesDocsForComponent(
		client,
		rootDir,
		"Gloo Mesh",
		GlooMeshRepoName,
		"v1.0.0-beta16",
		ossFileMapping,
	); err != nil {
		return err
	}

	// Gloo Mesh Enterprise
	if err := copyHelmValuesDocsForComponent(
		client,
		rootDir,
		"Gloo Mesh Enterprise",
		GlooMeshEnterpriseRepoName,
		"v1.0.0-beta16",
		enterpriseFileMapping,
	); err != nil {
		return err
	}

	// Gloo Mesh Enterprise RBAC Webhook
	if err := copyHelmValuesDocsForComponent(
		client,
		rootDir,
		"Enterprise RBAC Webhook",
		GlooMeshEnterpriseRepoName,
		"v1.0.0-beta16",
		rbacWwebookFileMapping,
	); err != nil {
		return err
	}

	return nil
}

// fetch Helm Values documentation from repo up to and including the version specified by earliestVerison
// fileMapping specifies a mapping from the file path in the origin repo to the file path in this repo
func copyHelmValuesDocsForComponent(
	client *github.Client,
	rootDir string,
	componentName string,
	repoName string,
	earliestVersion string,
	fileMapping map[string]string,
) error {
	// flush directory for idempotence
	helmDocsDir := filepath.Join(rootDir, helmDocsDir, strcase.SnakeCase(componentName))
	os.RemoveAll(helmDocsDir)
	os.MkdirAll(helmDocsDir, 0755)

	if err := createFileIfNotExists(helmDocsDir+"/"+"_index.md", fmt.Sprintf(helmValuesIndex, componentName)); err != nil {
		return eris.Errorf("error creating Helm values index file: %v", err)
	}

	// include Helm values docs for all versions > earliestVersion
	releases, _, err := client.Repositories.ListReleases(
		context.Background(),
		GithubOrg,
		repoName,
		&github.ListOptions{Page: 0, PerPage: 1000000},
	)
	if err != nil {
		return eris.Errorf("error listing releases: %v", err)
	}
	var tags []string
	for _, release := range releases {
		if release.GetTagName() == earliestVersion {
			break
		}
		tags = append(tags, release.GetTagName())
	}

	for _, tag := range tags {
		if err := os.Mkdir(helmDocsDir+"/"+tag, os.ModePerm); err != nil {
			return eris.Errorf("error creating Helm docs directories: %v", err)
		}

		if err := createFileIfNotExists(helmDocsDir+"/"+tag+"/"+"_index.md", fmt.Sprintf(helmValuesIndex, tag)); err != nil {
			return eris.Errorf("error creating Helm values index file: %v", err)
		}

		for src, dest := range fileMapping {
			dest = fmt.Sprintf(dest, helmDocsDir, tag)
			if err := copyHelmValuesDocs(client, GithubOrg, repoName, tag, src, dest); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyHelmValuesDocs(client *github.Client, org, repo, tag, path, destinationFile string) error {
	contents, _, resp, err := client.Repositories.GetContents(context.Background(), org, repo, path, &github.RepositoryContentGetOptions{
		Ref: tag,
	})
	if err != nil && resp.StatusCode != 404 {
		return eris.Errorf("error fetching Helm values doc: %v", err)
	} else if resp.StatusCode == 404 {
		contextutils.LoggerFrom(context.Background()).Warnf("Helm values doc for repo: %s and path: %s not found", repo, path)
		return nil
	}

	decodedContents, err := base64.StdEncoding.DecodeString(*contents.Content)
	if err != nil {
		return eris.Errorf("error fetching Helm values doc: %v", err)
	}

	return createFileIfNotExists(destinationFile, string(decodedContents))
}

// create file with contents, create the file if it doesn't exist
func createFileIfNotExists(fname, contents string) error {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return eris.Errorf("error copying Helm values reference doc: %v", err)
	}

	if _, err = f.Write([]byte(contents)); err != nil {
		return err
	}

	return nil
}
