import asyncio
import websockets
import json

CLIENTS = set()


async def notify(sender, message):
    clients = [c for c in CLIENTS if c != sender]
    for c in clients:
        await c.send(message)


async def register_client(ws):
    CLIENTS.add(ws)


async def unregister_client(ws):
    CLIENTS.remove(ws)


async def handler(ws, path):
    print("-- websocket connected --")
    await register_client(ws)
    try:
        async for message in ws:
            print(f"< {message}")
            await notify(ws, message)
    finally:
        await unregister_client(ws)
        print("- disconnected -")


start_server = websockets.serve(handler, None, 3001)

asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
