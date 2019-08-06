# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import api_pb2 as api__pb2


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
        request_serializer=api__pb2.CreateProjectRequest.SerializeToString,
        response_deserializer=api__pb2.CreateProjectResponse.FromString,
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


def add_CatalogServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateProject': grpc.unary_unary_rpc_method_handler(
          servicer.CreateProject,
          request_deserializer=api__pb2.CreateProjectRequest.FromString,
          response_serializer=api__pb2.CreateProjectResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Catalog', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class TestStub(object):
  """TODO: make this streaming

  Test service helps to automate specification checks, screenshots and tutorials
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Setup = channel.unary_unary(
        '/Test/Setup',
        request_serializer=api__pb2.ScenarioRequest.SerializeToString,
        response_deserializer=api__pb2.ScenarioResponse.FromString,
        )


class TestServicer(object):
  """TODO: make this streaming

  Test service helps to automate specification checks, screenshots and tutorials
  """

  def Setup(self, request, context):
    """Setup a given state in the database
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TestServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Setup': grpc.unary_unary_rpc_method_handler(
          servicer.Setup,
          request_deserializer=api__pb2.ScenarioRequest.FromString,
          response_serializer=api__pb2.ScenarioResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Test', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
