import time
from demo import analytics
from client import ml_pipelines as client


cl = client.connect('localhost:9111')

def wait_for_server_to_start():

    for i in range(100):
        try:
            cl.stats()
            return
        except client.Unavailable:
            time.sleep(0.1)
            continue
        except:
            print("breaking!!!!!!!!!!!!!!!!!!!!!!!!!")
            raise

    raise Exception("Failed to connect")


wait_for_server_to_start()

analytics.setup_analytics_demo(cl)

