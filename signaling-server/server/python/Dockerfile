FROM python:3.7-alpine

RUN apk add --update --no-cache --virtual build-deps \
    gcc \
    make \
    musl-dev \
    && pip install websockets \
    && apk del build-deps

COPY signaling.py /app/signaling.py

CMD ["python3", "/app/signaling.py"]
