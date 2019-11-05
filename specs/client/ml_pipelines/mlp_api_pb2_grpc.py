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
    self.ListProjects = channel.unary_unary(
        '/Catalog/ListProjects',
        request_serializer=mlp__api__pb2.ListProjectsRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.ListProjectsResponse.FromString,
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
    self.ListDatasets = channel.unary_unary(
        '/Catalog/ListDatasets',
        request_serializer=mlp__api__pb2.ListDatasetsRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.ListDatasetsResponse.FromString,
        )
    self.UpdateDataset = channel.unary_unary(
        '/Catalog/UpdateDataset',
        request_serializer=mlp__api__pb2.UpdateDatasetRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.EmptyResponse.FromString,
        )
    self.AddDatasetActivity = channel.unary_unary(
        '/Catalog/AddDatasetActivity',
        request_serializer=mlp__api__pb2.AddDatasetActivityRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.AddDatasetActivityResponse.FromString,
        )
    self.Stat = channel.unary_unary(
        '/Catalog/Stat',
        request_serializer=mlp__api__pb2.StatRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.StatResponse.FromString,
        )
    self.Reset = channel.unary_unary(
        '/Catalog/Reset',
        request_serializer=mlp__api__pb2.ResetRequest.SerializeToString,
        response_deserializer=mlp__api__pb2.EmptyResponse.FromString,
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

  def ListProjects(self, request, context):
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

  def ListDatasets(self, request, context):
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

  def AddDatasetActivity(self, request, context):
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

  def Reset(self, request, context):
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
      'ListProjects': grpc.unary_unary_rpc_method_handler(
          servicer.ListProjects,
          request_deserializer=mlp__api__pb2.ListProjectsRequest.FromString,
          response_serializer=mlp__api__pb2.ListProjectsResponse.SerializeToString,
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
      'ListDatasets': grpc.unary_unary_rpc_method_handler(
          servicer.ListDatasets,
          request_deserializer=mlp__api__pb2.ListDatasetsRequest.FromString,
          response_serializer=mlp__api__pb2.ListDatasetsResponse.SerializeToString,
      ),
      'UpdateDataset': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateDataset,
          request_deserializer=mlp__api__pb2.UpdateDatasetRequest.FromString,
          response_serializer=mlp__api__pb2.EmptyResponse.SerializeToString,
      ),
      'AddDatasetActivity': grpc.unary_unary_rpc_method_handler(
          servicer.AddDatasetActivity,
          request_deserializer=mlp__api__pb2.AddDatasetActivityRequest.FromString,
          response_serializer=mlp__api__pb2.AddDatasetActivityResponse.SerializeToString,
      ),
      'Stat': grpc.unary_unary_rpc_method_handler(
          servicer.Stat,
          request_deserializer=mlp__api__pb2.StatRequest.FromString,
          response_serializer=mlp__api__pb2.StatResponse.SerializeToString,
      ),
      'Reset': grpc.unary_unary_rpc_method_handler(
          servicer.Reset,
          request_deserializer=mlp__api__pb2.ResetRequest.FromString,
          response_serializer=mlp__api__pb2.EmptyResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Catalog', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
