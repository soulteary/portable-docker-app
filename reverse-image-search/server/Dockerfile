FROM python:3.9-slim-buster
RUN sed -i -E "s/\w+.debian.org/mirrors.tuna.tsinghua.edu.cn/g" /etc/apt/sources.list
RUN apt-get update && \
    apt-get install -y ffmpeg libsm6 libxext6 && \
    apt-get remove --purge -y && rm -rf /var/lib/apt/lists/*

RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
    pip install diskcache==5.2.1 fastapi==0.65.2 pymilvus==2.1.0 towhee==0.8.0 uvicorn==0.13.4 protobuf==3.20.2 opencv-python==4.6.0.66 torch==1.12.1 torchvision==0.13.1 Pillow==9.2.0 timm==0.6.7 && \
    pip cache purge

WORKDIR /app/src
COPY . /app

RUN python3 encode.py
CMD python3 main.py