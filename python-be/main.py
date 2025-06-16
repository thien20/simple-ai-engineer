from fastapi import FastAPI, HTTPException
from schemas import RAGRequest, RAGResponse
from handler.rag_handler import RagHandler
import uvicorn

app = FastAPI(title="Python RAG API", version="1.0")
ragHandler = RagHandler()

@app.post("/rag", response_model=RAGResponse)
async def rag(request: RAGRequest):
    """
    Handle RAG request.
    """
    try:
        return ragHandler.handle(request)
    except HTTPException as e:
        raise e
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    
if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8001, reload=True)