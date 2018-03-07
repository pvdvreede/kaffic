# kaffic

This repo is an experiment to create a messaging broker like Kafka or AWS Kinesis, but that is simple to setup and maintain. It trades high throughput for the simplicity and reliability.

The core value proposition for `kaffic` is to use a dumb, but highly available and reliable storage solution such as AWS S3 or Google Cloud Storage to write and read messages. This reduces the setup and maintenance by relying on the storage's high availability as well as ACL policies to restrict access and lifecycle policies to determine the retention of the topics. Of course this trades off the high throughput as the `kaffic` clients will have high latency talking to the storage layer.

`kaffic` is an experiment to see if that tradeoff is worthwhile for apps that need a reliable, highly available and easy to setup and maintain messaging system that can have near infinite message retention, and do NOT have any (or have lax) requirements around throughput.

## Objectives

* Provide 'Topics' that maintain FIFO and can be stored for potentially forever.
* Allow clients to choose where in the Topic they want to start consuming from.
* Enable scaling of a Topic using Kafka/Kinesis concepts like Shard (or Partition) each of which are FIFO based on a Partition Key.
* Library that can be used anywhere, so be able to compile this app as a C library that can be accessed with FFI in other languages such as Java, Python and Ruby.
* Pluggable storage to utilize dumb storage services like S3 or GCS as well as the FileSystem for development or simple Messaging requirements.
* Focus on reliability and easy setup/maintenance over throughput and latency.
* Defer ACL/Security/Lifecycle of data to underlying storage or by placing a HTTP proxy in front of the client.

`kaffic` is written in Go as it is reasonably fast, concurrent and can now create C-like shared libraries for interop with other languages. This will enable:

* Being able to add `kaffic` to a Go project that wants to use a messaging system as either a producer or consumer.
* Write a package for other languages using that language's FFI functionality. This is possible at least with Python, Ruby, NodeJS if not others.
* Be able to wrap this library as an HTTP API server for other services to consume that cannot use the library directly.

## Roadmap

[ ] Start with simple Go only package
[ ] Provide Write api for Topics (exclude scalable semantics like Paritions)
[ ] Start with a basic File System storage plugin to make it easy to test and confirm correctness
[ ] Add Read API for Topics (exclude scalable semantics like Paritions)
[ ] Add S3 pluggable storage to experiment with latency, throughput

If useful, further tasks would be:

[ ] Build in Go web server with HTTP API
[ ] Add in scaling semantics with partitions that can only be set at Topic creation
[ ] Make library shareable as a C Shared library
[ ] Create Python package with bindings to `kaffic`
[ ] Add GCS pluggable storage
[ ] Create Java package with bindings to `kaffic`
[ ] Create Ruby gem with bindings to `kaffic`
[ ] Create Node package with bindings to `kaffic`
