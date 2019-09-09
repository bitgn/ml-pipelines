# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import mlp_api_pb2 as mlp__api__pb2


class CatalogStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateProject = channel.unary_unary(
        '/Catalog/CreateProject',
        request_serializer=mlp__api__pb2.CreateProjectRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.ProjectInfoResponse.FromString,
        )
    self.GetProject = channel.unary_unary(
        '/Catalog/GetProject',
        request_serializer=mlp__api__pb2.GetProjectRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.ProjectInfoResponse.FromString,
        )
    self.AddSystem = channel.unary_unary(
        '/Catalog/AddSystem',
        request_serializer=mlp__api__pb2.AddSystemRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.SystemInfoResponse.FromString,
        )
    self.GetSystem = channel.unary_unary(
        '/Catalog/GetSystem',
        request_serializer=mlp__api__pb2.GetSystemRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.SystemInfoResponse.FromString,
        )
    self.AddSystemVersion = channel.unary_unary(
        '/Catalog/AddSystemVersion',
        request_serializer=mlp__api__pb2.AddSystemVersionRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.AddSystemVersionResponse.FromString,
        )
    self.CreateJob = channel.unary_unary(
        '/Catalog/CreateJob',
        request_serializer=mlp__api__pb2.CreateJobRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.JobInfoResponse.FromString,
        )
    self.GetJob = channel.unary_unary(
        '/Catalog/GetJob',
        request_serializer=mlp__api__pb2.GetJobRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.JobInfoResponse.FromString,
        )
    self.StartJobRun = channel.unary_unary(
        '/Catalog/StartJobRun',
        request_serializer=mlp__api__pb2.StartJobRunRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.JobRunInfoResponse.FromString,
        )
    self.LogJobRun = channel.unary_unary(
        '/Catalog/LogJobRun',
        request_serializer=mlp__api__pb2.LogJobRunRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.EmptyResponse.FromString,
        )
    self.FailJobRun = channel.unary_unary(
        '/Catalog/FailJobRun',
        request_serializer=mlp__api__pb2.FailJobRunRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.EmptyResponse.FromString,
        )
    self.CompleteJobRun = channel.unary_unary(
        '/Catalog/CompleteJobRun',
        request_serializer=mlp__api__pb2.CompleteJobRunRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.EmptyResponse.FromString,
        )
    self.CreateDataset = channel.unary_unary(
        '/Catalog/CreateDataset',
        request_serializer=mlp__api__pb2.CreateDatasetRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.DatasetInfoResponse.FromString,
        )
    self.GetDataset = channel.unary_unary(
        '/Catalog/GetDataset',
        request_serializer=mlp__api__pb2.GetDatasetRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.DatasetInfoResponse.FromString,
        )
    self.UpdateDataset = channel.unary_unary(
        '/Catalog/UpdateDataset',
        request_serializer=mlp__api__pb2.UpdateDatasetRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.EmptyResponse.FromString,
        )
    self.GetLastDatasetVersion = channel.unary_unary(
        '/Catalog/GetLastDatasetVersion',
        request_serializer=mlp__api__pb2.GetLastDatasetVersionRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.DatasetVersionResponse.FromString,
        )
    self.Commit = channel.unary_unary(
        '/Catalog/Commit',
        request_serializer=mlp__api__pb2.CommitRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.CommitResponse.FromString,
        )
    self.Stat = channel.unary_unary(
        '/Catalog/Stat',
        request_serializer=mlp__api__pb2.StatRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.StatResponse.FromString,
        )


class CatalogServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def CreateProject(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetProject(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AddSystem(self, request, context):
    """SYSTEMS ------------

    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetSystem(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AddSystemVersion(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateJob(self, request, context):
    """JOBS -----------------
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetJob(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StartJobRun(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def LogJobRun(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FailJobRun(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CompleteJobRun(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateDataset(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetDataset(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateDataset(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetLastDatasetVersion(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Commit(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Stat(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CatalogServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateProject': grpc.unary_unary_rpc_method_handler(
          servicer.CreateProject,
          request_deserializer=mlp__api__pb2.CreateProjectRequest.FromString,
          response_serializer=mlp__api__pb2.ProjectInfoResponse.SerializeToString,
      ),
      'GetProject': grpc.unary_unary_rpc_method_handler(
          servicer.GetProject,
          request_deserializer=mlp__api__pb2.GetProjectRequest.FromString,
          response_serializer=mlp__api__pb2.ProjectInfoResponse.SerializeToString,
      ),
      'AddSystem': grpc.unary_unary_rpc_method_handler(
          servicer.AddSystem,
          request_deserializer=mlp__api__pb2.AddSystemRequest.FromString,
          response_serializer=mlp__api__pb2.SystemInfoResponse.SerializeToString,
      ),
      'GetSystem': grpc.unary_unary_rpc_method_handler(
          servicer.GetSystem,
          request_deserializer=mlp__api__pb2.GetSystemRequest.FromString,
          response_serializer=mlp__api__pb2.SystemInfoResponse.SerializeToString,
      ),
      'AddSystemVersion': grpc.unary_unary_rpc_method_handler(
          servicer.AddSystemVersion,
          request_deserializer=mlp__api__pb2.AddSystemVersionRequest.FromString,
          response_serializer=mlp__api__pb2.AddSystemVersionResponse.SerializeToString,
      ),
      'CreateJob': grpc.unary_unary_rpc_method_handler(
          servicer.CreateJob,
          request_deserializer=mlp__api__pb2.CreateJobRequest.FromString,
          response_serializer=mlp__api__pb2.JobInfoResponse.SerializeToString,
      ),
      'GetJob': grpc.unary_unary_rpc_method_handler(
          servicer.GetJob,
          request_deserializer=mlp__api__pb2.GetJobRequest.FromString,
          response_serializer=mlp__api__pb2.JobInfoResponse.SerializeToString,
      ),
      'StartJobRun': grpc.unary_unary_rpc_method_handler(
          servicer.StartJobRun,
          request_deserializer=mlp__api__pb2.StartJobRunRequest.FromString,
          response_serializer=mlp__api__pb2.JobRunInfoResponse.SerializeToString,
      ),
      'LogJobRun': grpc.unary_unary_rpc_method_handler(
          servicer.LogJobRun,
          request_deserializer=mlp__api__pb2.LogJobRunRequest.FromString,
          response_serializer=mlp__api__pb2.EmptyResponse.SerializeToString,
      ),
      'FailJobRun': grpc.unary_unary_rpc_method_handler(
          servicer.FailJobRun,
          request_deserializer=mlp__api__pb2.FailJobRunRequest.FromString,
          response_serializer=mlp__api__pb2.EmptyResponse.SerializeToString,
      ),
      'CompleteJobRun': grpc.unary_unary_rpc_method_handler(
          servicer.CompleteJobRun,
          request_deserializer=mlp__api__pb2.CompleteJobRunRequest.FromString,
          response_serializer=mlp__api__pb2.EmptyResponse.SerializeToString,
      ),
      'CreateDataset': grpc.unary_unary_rpc_method_handler(
          servicer.CreateDataset,
          request_deserializer=mlp__api__pb2.CreateDatasetRequest.FromString,
          response_serializer=mlp__api__pb2.DatasetInfoResponse.SerializeToString,
      ),
      'GetDataset': grpc.unary_unary_rpc_method_handler(
          servicer.GetDataset,
          request_deserializer=mlp__api__pb2.GetDatasetRequest.FromString,
          response_serializer=mlp__api__pb2.DatasetInfoResponse.SerializeToString,
      ),
      'UpdateDataset': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateDataset,
          request_deserializer=mlp__api__pb2.UpdateDatasetRequest.FromString,
          response_serializer=mlp__api__pb2.EmptyResponse.SerializeToString,
      ),
      'GetLastDatasetVersion': grpc.unary_unary_rpc_method_handler(
          servicer.GetLastDatasetVersion,
          request_deserializer=mlp__api__pb2.GetLastDatasetVersionRequest.FromString,
          response_serializer=mlp__api__pb2.DatasetVersionResponse.SerializeToString,
      ),
      'Commit': grpc.unary_unary_rpc_method_handler(
          servicer.Commit,
          request_deserializer=mlp__api__pb2.CommitRequest.FromString,
          response_serializer=mlp__api__pb2.CommitResponse.SerializeToString,
      ),
      'Stat': grpc.unary_unary_rpc_method_handler(
          servicer.Stat,
          request_deserializer=mlp__api__pb2.StatRequest.FromString,
          response_serializer=mlp__api__pb2.StatResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Catalog', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
