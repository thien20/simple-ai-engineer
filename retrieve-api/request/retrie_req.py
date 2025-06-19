from pydantic import BaseModel

class RetrieveRequest(BaseModel):
    userInput: str
    top_k: int = 1