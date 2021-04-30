---
title: "Install Gloo Mesh Istio"
menuTitle: Install Gloo Mesh Istio
description: Installing Gloo Mesh Istio
weight: 10
---

{{% notice note %}}
Gloo Mesh Enterprise is required for this feature. Open source users and users who do not require the functionality
provided by Gloo Mesh Istio can use Gloo Mesh with an upstream Istio release and should refer to our guide on [installing Istio]({{% versioned_link_path fromRoot="/guides/installing_istio" %}}).
{{% /notice %}}

Gloo Mesh Istio consists of upstream builds of both the Istio control plane and Istio data plane. These builds are available to
enterprise users and can contain patches, security back-ports and other fixes. For example, a Gloo Mesh Istio build of 1.7.5 (out of support from community) can contain security back-ports for CVEs found in 1.8/1.9 (current community support) or feature backports for newer versions. Any back-ported functionality is all available in the upstream in some version -- there are no proprietary features or forked capabilities from upstream. It's all based on upstream in the support of our enterprise users. 

To install Gloo Mesh Istio, simply override the `tag` and `hub` parameters on the [Istio Operator](https://istio.io/latest/docs/reference/config/istio.operator.v1alpha1/)
resource at install time. The `hub` value must equal `gcr.io/istio-enterprise` and the `tag` value should equal the
desired tag of Gloo Mesh Istio. For example, use tag `1.7.5-fips2` to install Gloo Mesh Istio builds that are compliant
with Federal Information Processing Standards.

The following install commands are lifted from the [Installing Istio Multicluster guide]({{% versioned_link_path fromRoot="/guides/installing_istio" %}}),
but the `hub` and `tag` values can be added to any Istio installation manifest to install Gloo Mesh Istio.

{{< tabs >}}
{{< tab name="Istio 1.9" codelang="shell" >}}
cat << EOF | istioctl manifest install -y -f -
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  name: gloo-mesh-istio
  namespace: istio-system
spec:
  # This value is required for Gloo Mesh Istio
  hub: gcr.io/istio-enterprise
  # This value can be any Gloo Mesh Istio tag
  tag: 1.9.2
  profile: minimal
  meshConfig:
    enableAutoMtls: true
    defaultConfig:
      proxyMetadata:
        # Enable Istio agent to handle DNS requests for known hosts
        # Unknown hosts will automatically be resolved using upstream dns servers in resolv.conf
        ISTIO_META_DNS_CAPTURE: "true"
  components:
    # Istio Gateway feature
    ingressGateways:
    - name: istio-ingressgateway
      enabled: true
      k8s:
        env:
          - name: ISTIO_META_ROUTER_MODE
            value: "sni-dnat"
        service:
          type: ClusterIP
          ports:
            - port: 80
              targetPort: 8080
              name: http2
            - port: 443
              targetPort: 8443
              name: https
            - port: 15443
              targetPort: 15443
              name: tls
  values:
    global:
      pilotCertProvider: istiod
EOF
{{< /tab >}}
{{< tab name="Istio 1.8" codelang="shell" >}}
cat << EOF | istioctl manifest install -y -f -
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  name: gloo-mesh-istio
  namespace: istio-system
spec:
  # This value is required for Gloo Mesh Istio
  hub: gcr.io/istio-enterprise
  # This value can be any Gloo Mesh Istio tag
  tag: 1.8.4
  profile: minimal
  meshConfig:
    enableAutoMtls: true
    defaultConfig:
      proxyMetadata:
        # Enable Istio agent to handle DNS requests for known hosts
        # Unknown hosts will automatically be resolved using upstream dns servers in resolv.conf
        ISTIO_META_DNS_CAPTURE: "true"
  components:
    # Istio Gateway feature
    ingressGateways:
    - name: istio-ingressgateway
      enabled: true
      k8s:
        env:
          - name: ISTIO_META_ROUTER_MODE
            value: "sni-dnat"
        service:
          type: ClusterIP
          ports:
            - port: 80
              targetPort: 8080
              name: http2
            - port: 443
              targetPort: 8443
              name: https
            - port: 15443
              targetPort: 15443
              name: tls
              nodePort: 32000
  values:
    global:
      pilotCertProvider: istiod
EOF
{{< /tab >}}
{{< tab name="Istio 1.7" codelang="shell" >}}
cat << EOF | istioctl manifest install -f -
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  name: gloo-mesh-istio
  namespace: istio-system
spec:
  # This value is required for Gloo Mesh Istio
  hub: gcr.io/istio-enterprise
  # This value can be any Gloo Mesh Istio tag
  tag: 1.7.5-fips2
  profile: minimal
  addonComponents:
    istiocoredns:
      enabled: true
  components:
    # Istio Gateway feature
    ingressGateways:
      - name: istio-ingressgateway
        enabled: true
        k8s:
          env:
            - name: ISTIO_META_ROUTER_MODE
              value: "sni-dnat"
          service:
            ports:
              - port: 80
                targetPort: 8080
                name: http2
              - port: 443
                targetPort: 8443
                name: https
              - port: 15443
                targetPort: 15443
                name: tls
                nodePort: 32000
  meshConfig:
    enableAutoMtls: true
  values:
    prometheus:
      enabled: false
    gateways:
      istio-ingressgateway:
        type: NodePort
        ports:
          - targetPort: 15443
            name: tls
            nodePort: 32000
            port: 15443
    global:
      pilotCertProvider: istiod
      controlPlaneSecurityEnabled: true
      podDNSSearchNamespaces:
        - global

EOF
{{< /tab >}}
{{< /tabs >}}

After installation is complete, you should see that istiod and an ingress gateway pods have been created with the Gloo Mesh images.

```shell
kubectl get pods -n istio-system -o=jsonpath='{range .items[*]}{"\n"}{.metadata.name}{":\t"}{range .spec.containers[*]}{.image}{", "}{end}{end}'
```

```shell
istio-ingressgateway-ff54575c8-97c67:	gcr.io/istio-enterprise/proxyv2:1.8.4,
istiod-5fbb57854-9pl28:	gcr.io/istio-enterprise/pilot:1.8.4,
```

## Next steps

Now that we have Istio and Gloo Mesh installed ([and appropriate clusters registered]({{% versioned_link_path fromRoot="/setup/#register-a-cluster" %}})), we can continue to explore the [discovery capabilities]({{% versioned_link_path fromRoot="/guides/discovery_intro" %}}) of Gloo Mesh. 
