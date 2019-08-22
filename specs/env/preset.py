from google.protobuf.internal import well_known_types as wkt

from test_api import events_pb2 as evt
from test_api import vo_pb2 as vo
import test
import json
import env
import random
from datetime import timedelta

import faker

nouns = ['data', 'trends', 'analytics', 'dump', 'dataset', 'table', 'stream', 'model', 'forecast']

verbs = ['research', 'demo', 'explore', 'examine', 'mine', 'analyze', 'investigate', 'clean', 'improve']

adj = ['learning', 'deep', 'derived', 'bot', 'iot', 'sensor', 'cargo', 'sales', 'billing', 'inventory', 'po', 'market']




def project_created(e: env.Env):
    id = e.next_id()

    n = random.choice(nouns)
    v = random.choice(verbs)
    a = random.choice(adj)

    name = f'{v.title()} {a} {n} prj{id}'

    pid = f'{v}-{a}-{n}-{id}'

    d = vo.ProjectMetadataDelta(
        name=name,
        name_set=True,
    )

    return evt.ProjectCreated(project_id=pid, meta=d)




def set_update_timestamp(t: env.Env, dm: vo.DatasetMetadataDelta, days=0, minutes = 0):
    time = t.time + timedelta(days=days, minutes=minutes)
    dm.update_timestamp = int(time.timestamp())
    dm.update_timestamp_set = True

def set_sample(t: env.Env, dm: vo.DatasetMetadataDelta, text: str, type=vo.DatasetSample.FORMAT.TEXT):
    dm.sample.body = text.encode()
    dm.sample.format = type
    dm.sample_set = True


def set_description(t: env.Env, dm: vo.DatasetMetadataDelta, text: str):
    dm.description = text.encode()
    dm.description_set = True

def set_data_format(t: env.Env, dm:vo.DatasetMetadataDelta, text: str):
    dm.data_format = text
    dm.data_format_set = True

def dataset_created(t: env.Env, e: evt.ProjectCreated):
    id = t.next_id()
    a = random.choice(adj)
    n = random.choice(nouns)

    name = f'{a} {n} ds{id}'.title()
    did = f'{id:08x}'



    meta = vo.DatasetMetadataDelta()

    meta.name = name
    meta.name_set = True

    meta.record_count = random.randint(10, 3000)
    meta.record_count_set = True

    raw_bytes = random.randint(100, 2 ** 30)
    meta.storage_bytes = random.randint(int(raw_bytes / 2), int(raw_bytes))
    meta.storage_bytes_set = True

    meta.file_count = random.randint(1, meta.record_count)
    meta.file_count_set = True


    set_update_timestamp(t, meta, minutes=-random.randint(0, 7 * 24 * 60))



    kind = random.randint(0, 3)
    archive = random.randint(0, 1) == 1

    if kind == 1:
        text = random.choice(json_samples)
        text = json.dumps(json.loads(text), indent=2)
        set_sample(t, meta, text, vo.DatasetSample.FORMAT.JSON)

        if archive:
            meta.data_format = "json+gzip"
        else:
            meta.data_format = "json"
        meta.data_format_set = True


    if random.randint(0,2)==1:
        meta.description_set = True
        meta.description = """
# Sales
**Sales** are activities related to selling or the number of goods or services sold in a given time period.

The _seller_, or the provider of the goods or services, completes a sale in response to an acquisition, appropriation,[1] requisition, or a direct interaction with the buyer at the point of sale. There is a passing of title (property or ownership) of the item, and the settlement of a price, in which agreement is reached on a price for which transfer of ownership of the item will occur. The seller, not the purchaser, typically executes the sale and it may be completed prior to the obligation of payment. In the case of indirect interaction, a person who sells goods or service on behalf of the owner is known as a salesman or saleswoman or salesperson, but this often refers to someone selling goods in a store/shop, in which case other terms are also common, including salesclerk, shop assistant, and retail clerk.        
        """.strip()

    return evt.DatasetCreated(
        dataset_id=did,
        project_id=e.project_id,
        meta=meta
    )

def job_added(e:env.Env, prj:evt.ProjectCreated) -> evt.JobAdded:
    id  = e.next_id()
    jid = f'job-{id}'

    name = f'Job {id}'
    return evt.JobAdded(job_id=jid, job_name=name, project_id=prj.project_id)

def expert_added(e:env.Env) -> evt.ExpertAdded:

    id = f'expert-{e.next_id()}'

    return evt.ExpertAdded(expert_id=id, expert_name=e.faker.name())




json_samples = [
    """{"lat":"39.940631","lng":"116.536442","gps_type":"baidu","gps_time":1418779410000,"provider":"gps","accuracy":2,"status":1,"milli_timestamp":1418779410.136}""",
    """{"accuracy": 61.581451, "gps_time": 1418777594000.000000, "gps_type": "baidu", "lat": "39.919616", "lng": "116.585133", "milli_timestamp": 1418777593.054000, "provider": "network", "status": 1}"""

    ,
    """{ "name"   : "John Smith",
  "sku"    : "20223",
  "price"  : 23.95,
  "shipTo" : { "name" : "Jane Smith",
               "address" : "123 Maple Street",
               "city" : "Pretendville",
               "state" : "NY",
               "zip"   : "12345" },
  "billTo" : { "name" : "John Smith",
               "address" : "123 Maple Street",
               "city" : "Pretendville",
               "state" : "NY",
               "zip"   : "12345" }
}""",
    """{
  "redis": {
    "host": "mint.vm.ge.rgunti.ch",
    "port": 6379,
    "keys": {
      "gpsAlive": "DaemonAlive.GPS",
      "obdAlive": "OBD.State",
      "latitude": "GPS.Latitude",
      "longitude": "GPS.Longitude",
      "altitude": "GPS.Altitude",
      "fixMode": "GPS.FixMode",
      "gpsTime": "GPS.Time",
      "gpsSpeed": "GPS.Speed",
      "gpsSpeedKMH": "GPS.Speed.KMH",
      "gpsAccuracyLon": "GPS.EPX",
      "gpsAccuracyLat": "GPS.EPY",
      "gpsAccuracyHeight": "GPS.EPV",
      "gpsAccuracySpeed": "GPS.EPS",
      "obdSpeedKMH": "OBD.Speed",
      "engineRpm": "OBD.RPM",
      "intakeTemp": "OBD.IntakeTemp",
      "intakeMAP": "OBD.IntakeMAP",
      "recordTripIDA": "Trip.A.ID",
      "recordTripIDB": "Trip.B.ID"
    }
  },
  "persistentRedis": {
    "host": "mint.vm.ge.rgunti.ch",
    "port": 6379,
    "keys": {
      "odo": "GPS.ODO",
      "tripA": "GPS.Trip.A",
      "tripB": "GPS.Trip.B"
    }
  },
  "database": {
    "host": "192.168.1.250",
    "port": 3306,
    "user": "carpi-rec",
    "pass": "carpi-rec",
    "db": "carpi-rec"
  },
  "interval": 250,
  "minTravelDistanceM": "10"
}""", """ {
      "type": "Payment",
      "id": "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
      "version": 0,
      "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
      "attributes": {
        "amount": "100.21",
        "beneficiary_party": {
          "account_name": "W Owens",
          "account_number": "31926819",
          "account_number_code": "BBAN",
          "account_type": 0,
          "address": "1 The Beneficiary Localtown SE2",
          "bank_id": "403000",
          "bank_id_code": "GBDSC",
          "name": "Wilfred Jeremiah Owens"
        },
        "charges_information": {
          "bearer_code": "SHAR",
          "sender_charges": [
            {
              "amount": "5.00",
              "currency": "GBP"
            },
            {
              "amount": "10.00",
              "currency": "USD"
            }
          ],
          "receiver_charges_amount": "1.00",
          "receiver_charges_currency": "USD"
        },
        "currency": "GBP",
        "debtor_party": {
          "account_name": "EJ Brown Black",
          "account_number": "GB29XABC10161234567801",
          "account_number_code": "IBAN",
          "address": "10 Debtor Crescent Sourcetown NE1",
          "bank_id": "203301",
          "bank_id_code": "GBDSC",
          "name": "Emelia Jane Brown"
        },
        "end_to_end_reference": "Wil piano Jan",
        "fx": {
          "contract_reference": "FX123",
          "exchange_rate": "2.00000",
          "original_amount": "200.42",
          "original_currency": "USD"
        },
        "numeric_reference": "1002001",
        "payment_id": "123456789012345678",
        "payment_purpose": "Paying for goods/services",
        "payment_scheme": "FPS",
        "payment_type": "Credit",
        "processing_date": "2017-01-18",
        "reference": "Payment for Em's piano lessons",
        "scheme_payment_sub_type": "InternetBanking",
        "scheme_payment_type": "ImmediatePayment",
        "sponsor_party": {
          "account_number": "56781234",
          "bank_id": "123123",
          "bank_id_code": "GBDSC"
        }
      }
    }"""]
