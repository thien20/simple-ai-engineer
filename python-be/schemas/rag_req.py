from pydantic import BaseModel

class RAGRequest(BaseModel):
    userInput: str