from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
import random
import time

app = FastAPI()

@app.get("/py/hello", response_class=PlainTextResponse)
async def hello_handler():
    sleep_time = random.uniform(0, 5)
    time.sleep(sleep_time)
    return "Hello, Python\n"

if __name__ == "__main__":
    import uvicorn
    print(f"python server is listening on :8000...")
    uvicorn.run(app, host="127.0.0.1", port=8000)
