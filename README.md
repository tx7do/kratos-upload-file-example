# Kratos 上传文件实例

本实例演示了如何使用Kratos上传文件。上传的文件都落地于MinIO。

演示了后端的两种上传方式：

1. 通过像Kratos的服务申请预签名URL，然后通过预签名URL向MinIO上传文件。
2. 直接向Kratos的服务上传文件，然后，微服务再将文件落地到MinIO。

方式一是最优的解决方案，因为文件不会经过微服务，直接上传到MinIO，减轻了微服务的压力。并且，MinIO支持分布式部署，可以很好的扩展。

方式二是最简单的解决方案，但是不推荐，因为文件会经过微服务，增加了微服务的压力。
