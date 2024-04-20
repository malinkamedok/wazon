FROM python

COPY requirements.txt /requirements/requirements.txt
COPY . /app
RUN pip install -r /requirements/requirements.txt

CMD ["python", "/app/main.py"]