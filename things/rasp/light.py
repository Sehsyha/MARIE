import paho.mqtt.client as mqtt
from gpiozero import LED
from time import sleep
from utils import isAction, isGetter, register

led = LED(17)

on = 0

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    global on
    if isAction(msg.topic, "on"):
      on = 1
    else:
      on = 0
    print(msg.topic + " " + str(msg.payload))


mqttc = mqtt.Client()
mqttc.on_message = on_message
mqttc.on_connect = on_connect
mqttc.connect("10.0.10.3", 1883, 60)

actions = [
  { "name": "on"},
  { "name": "off"}
]
getters = []

register(mqttc, "Lumiere", "light", "couloir", actions, getters)


rc = 0
while rc == 0:
  if on == 1:
    led.on()
  if on == 0:
    led.off()
  rc = mqttc.loop()
