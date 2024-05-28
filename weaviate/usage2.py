from sentence_transformers import SentenceTransformer
from sentence_transformers.util import cos_sim

sentences = [
    "天气好热，哪里有卖冰棍的",
    "今天好冷，该多穿两件",
    "夏天",
    "冬天"
]

model_id = "./thenlper/gte-small"

model = SentenceTransformer(model_id)
embeddings = model.encode(sentences)
print(cos_sim(embeddings[0], embeddings[1]))
