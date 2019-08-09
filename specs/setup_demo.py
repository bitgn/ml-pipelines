import grpc

from api import api_pb2_grpc, marshal
from api import api_pb2 as api

from demo import analytics

channel = grpc.insecure_channel('localhost:50051')
stub = api_pb2_grpc.TestStub(channel)

events = list(analytics.setup_analytics_demo())
pack = marshal.serialize(events)
stub.Setup(api.ScenarioRequest(Events=pack))






