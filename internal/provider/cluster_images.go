package provider

// https://github.com/kubernetes-sigs/kind/releases/tag/v0.20.0
var clusterImage = map[string]string{
	"1.29.2":  "kindest/node:v1.29.2@sha256:51a1434a5397193442f0be2a297b488b6c919ce8a3931be0ce822606ea5ca245",
	"1.29.1":  "kindest/node:v1.29.1@sha256:0c06baa545c3bb3fbd4828eb49b8b805f6788e18ce67bff34706ffa91866558b",
	"1.28.7":  "kindest/node:v1.28.7@sha256:9bc6c451a289cf96ad0bbaf33d416901de6fd632415b076ab05f5fa7e4f65c58",
	"1.28.6":  "kindest/node:v1.28.6@sha256:e9e59d321795595d0eed0de48ef9fbda50388dc8bd4a9b23fb9bd869f370ec7e",
	"1.27.11": "kindest/node:v1.27.11@sha256:681253009e68069b8e01aad36a1e0fa8cf18bb0ab3e5c4069b2e65cafdd70843",
	"1.27.10": "kindest/node:v1.27.10@sha256:e6b2f72f22a4de7b957cd5541e519a8bef3bae7261dd30c6df34cd9bdd3f8476",
	"1.26.14": "kindest/node:v1.26.14@sha256:5d548739ddef37b9318c70cb977f57bf3e5015e4552be4e27e57280a8cbb8e4f",
	"1.26.13": "kindest/node:v1.26.13@sha256:8cb4239d64ff897e0c21ad19fe1d68c3422d4f3c1c1a734b7ab9ccc76c549605",
	"1.25.16": "kindest/node:v1.25.16@sha256:e8b50f8e06b44bb65a93678a65a26248fae585b3d3c2a669e5ca6c90c69dc519",
	"1.24.17": "kindest/node:v1.24.17@sha256:bad10f9b98d54586cba05a7eaa1b61c6b90bfc4ee174fdc43a7b75ca75c95e51",
	"1.23.17": "kindest/node:v1.23.17@sha256:14d0a9a892b943866d7e6be119a06871291c517d279aedb816a4b4bc0ec0a5b3",
}
