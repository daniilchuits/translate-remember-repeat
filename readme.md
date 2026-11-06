# Translate-remember-repeat

You write in CLI a word to translate, write a sentence 
in wich this word was used. First the program addresses 
to Redis, if it finds word here, it gives words translation
from Redis and updates it's ttl to 1 day. Then, if word was
unfound in Redis, service adresses word to Postgres db, if it 
existed in Postgres, the word with it translation and context
are recorded to Redis. If it does not exist - the program sends
request to TranslateAPI (don't know exact service) with context.
Service remembers this data via postgresql note: 
        (word; translation; context)
---
### **Service that I do to practise Redis a little, and it is some kind of usefull**

