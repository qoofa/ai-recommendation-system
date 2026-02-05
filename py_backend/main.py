from fastapi import FastAPI
from pydantic import BaseModel
from sentence_transformers import SentenceTransformer

app = FastAPI()

model = SentenceTransformer(
    "jinaai/jina-embeddings-v2-base-en",
    trust_remote_code=True
)

print("Model loaded successfully!")

class RequestText(BaseModel):
    text: str

@app.post("/embed")
async def embed_text(data: RequestText):
    embedding = model.encode(
        data.text,
        normalize_embeddings=True
    )
    return {
        "status": True,
        "embedding": embedding.tolist()
    }
