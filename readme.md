# Translate-remember-repeat

You write a word to translate in the CLI and write a sentence in which this word was used. First, the program addresses Redis. If it finds the word there, it returns the word's translation from Redis and updates its TTL to 1 day. Then, if the word is not found in Redis, the service addresses the PostgreSQL database. If it exists in PostgreSQL, the word, its translation, and context are recorded in Redis. If it does not exist, the program sends a request to the Translate API (don't know the exact service) with the context. The service remembers this data in PostgreSQL:
        (word; translation; context)

Found a library to easily deal with DeepL (deepL-go/deepL), but I wanted to do this manually.

LOOOOOOOOOOOL, implement API requests manually, but in the end, I realised that resp.StatusCode == 451 (Unavailable For Legal Reasons).

So, the solutions are to either use a VPN every time I run the code or use this library :D and a VPN. I'll use a VPN + my own implementation.

OK, it will be impossible to use if I create a container for the code. I'll run PostgreSQL and Redis in containers, but the code itself will run locally (using these services).

---

### **The service that I'm making to practise Redis a little, and it's also somewhat useful.**

### **Workflow of this translator:**

1. It invites you to write a sentence in the CLI that will be translated into English.
2. First, it looks for a translation in Redis. If the translation is found, the key's TTL is updated to 24 hours.
3. If it's not found in Redis, the program checks the PostgreSQL table to see if the sentence exists there. If the sentence is found in PostgreSQL, it creates a note with `sentence:(translation, detached_language)` in Redis with a TTL of 24 hours.
4. If the sentence is not found in Redis or PostgreSQL, the program sends an API request to translate it.
5. The user can also type "test" in the invitation zone to start a simple test with random words. The test consists of `domain.TestSentences` tasks. Each sentence waits for the user's translation, then notifies the user whether their answer was right or wrong. At the end of the test, the user gets their final score and can request an analysis of it (work in progress).

Also, apart from the "test" command, there will be "exit", "help", and some others (haven't come up with anything yet). First, I want to implement the analysis of the final results through RabbitMQ.
---
### Locally I did fast-forward merges, so the graph is not as nice as I thought.