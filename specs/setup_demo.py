import time

import grpc

from test_api import api_pb2_grpc, marshal
from test_api import api_pb2 as api

from demo import analytics

channel = grpc.insecure_channel('localhost:50051')
stub = api_pb2_grpc.TestStub(channel)

events = list(analytics.setup_analytics_demo())
pack = marshal.serialize(events)


def wait_for_server_to_start():
    while True:
        try:
            stub.Ping(api.PingRequest())
            break
        except:
            time.sleep(0.1)
            continue


wait_for_server_to_start()
stub.Setup(api.ScenarioRequest(Events=pack))






