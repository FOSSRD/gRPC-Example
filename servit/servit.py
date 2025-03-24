"""The Python implementation of the GRPC helloworld.Greeter client."""

from __future__ import print_function

import logging

import grpc
import converter_pb2
import converter_pb2_grpc
import streamlit as st
import cv2
import numpy as np


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    data = []
    with open("img.png", "rb") as image:
        f = image.read()
        data = f
    print("Will try to greet world ...")
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = converter_pb2_grpc.ImageConverterStub(channel)
        imagen = converter_pb2.Image(data=data, format="png")
        response1 = stub.BlackAndWhite(converter_pb2.BlackAndWhiteRequest(image=imagen,threshold=0.45))
        response2 = stub.Blur(converter_pb2.BlurRequest(image=imagen, kernel_size=30))
        # stub = helloworld_pb2_grpc.GreeterStub(channel)
        response3 = stub.Sepia(converter_pb2.SepiaRequest(image=imagen,intensity=1))
        # response = stub.SayHello(helloworld_pb2.HelloRequest(name="you"))
        return [response1.data, response2.data, response3.data]
        # print("Greeter client received: " + response.data)


# if __name__ == "__main__":
    # logging.basicConfig()
    # run()
photos = run()
print("fuimono")
file_bytes = np.asarray(bytearray(photos[0]), dtype=np.uint8)
opencv_image = cv2.imdecode(file_bytes, 1)
st.image(opencv_image, channels="BGR")
file_bytes = np.asarray(bytearray(photos[1]), dtype=np.uint8)
opencv_image = cv2.imdecode(file_bytes, 1)
st.image(opencv_image, channels="BGR")
file_bytes = np.asarray(bytearray(photos[2]), dtype=np.uint8)
opencv_image = cv2.imdecode(file_bytes, 1)
st.image(opencv_image, channels="BGR")
