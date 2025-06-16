from fastapi import HTTPException
from schemas import RAGRequest, RAGResponse
from constant.const import SystemPrompt, RetrieveApi
import requests

class RagHandler:
    def __init__(self):
        pass
    def handle(self, request: RAGRequest) -> RAGResponse:
        try:
            retriever_resp = requests.post(
                RetrieveApi,
                json={"query": request.userInput}
            )
            retriever_resp.raise_for_status()
            context = retriever_resp.json().get("context", "")

            prompt = f"{SystemPrompt}: {context}\n\nQuestion: {request.userInput}\nAnswer:"
            ollama_resp = requests.post(
                "http://ollama:11434/api/generate",
                json={
                    "model": "gemma",
                    "prompt": prompt,
                    "stream": False
                }
            )
            ollama_resp.raise_for_status()
            answer = ollama_resp.json().get("response", "")

            return RAGResponse(answer=answer, status=200)

        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

        
    