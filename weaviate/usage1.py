import torch.nn.functional as F
from transformers import AutoTokenizer, AutoModel

input_texts = [
    "天气好热，哪里有卖冰棍的",
    "今天好冷，该多穿两件",
    "夏天",
    "冬天"
]

# "thenlper/gte-base-zh"
model_id = "./thenlper/gte-small"

tokenizer = AutoTokenizer.from_pretrained(model_id)
model = AutoModel.from_pretrained(model_id)

# Tokenize the input texts
batch_dict = tokenizer(input_texts, max_length=512, padding=True, truncation=True, return_tensors='pt')

outputs = model(**batch_dict)
embeddings = outputs.last_hidden_state[:, 0]
 
# (Optionally) normalize embeddings
embeddings = F.normalize(embeddings, p=2, dim=1)
scores = (embeddings[:1] @ embeddings[1:].T) * 100
print(scores.tolist())

