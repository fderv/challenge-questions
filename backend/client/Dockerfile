FROM python:3.6
WORKDIR /client
ADD client.py requirements.txt user_pb2.py user_pb2_grpc.py /client/
ADD users /client/users
RUN pip install -r requirements.txt
RUN rm requirements.txt
CMD ["python", "client.py"]