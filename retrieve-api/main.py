from fastapi import FastAPI, HTTPException
from sentence_transformers import SentenceTransformer
from request.retrie_req import RetrieveRequest
import uvicorn
from elasticsearch import Elasticsearch

ELASTIC_INDEX = "knowledge_base"
es = Elasticsearch("http://localhost:9200")
model = SentenceTransformer("all-MiniLM-L6-v2")

app = FastAPI()

@app.post("/retrieve")
async def retrieve(request: RetrieveRequest):
    try:
        query_emb = model.encode([request.userInput]).tolist()[0]
        script_query = {
            "script_score": {
                "query": {"match_all": {}},
                "script": {
                    "source": "cosineSimilarity(params.query_vector, 'embedding') + 1.0",
                    "params": {"query_vector": query_emb}
                }
            }
        }
        res = es.search(
            index=ELASTIC_INDEX,
            body={"size": request.top_k, "query": script_query}
        )
        docs = [hit["_source"]["text"] for hit in res["hits"]["hits"]]
        return {"documents": docs}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=5002, reload=True)