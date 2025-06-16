from pydantic import BaseModel

class RAGResponse(BaseModel):
    result: str