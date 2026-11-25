#!/usr/bin/env bash
# Ask Witness 1 a question, with the test sources pasted into the prompt.
# Usage: ./test/ask-witness1.sh "Ki moun ki te Toussaint Louverture?"
# Requires: jq, and mlx_lm.server running on localhost:8080.

set -euo pipefail
cd "$(dirname "$0")/.."

QUESTION="${1:?Usage: $0 \"your question\"}"
URL="${LLM_URL:-http://localhost:8080/v1/chat/completions}"

jq -n \
  --rawfile system prompts/witness-1.md \
  --rawfile sources test/sources-w1.xml \
  --arg question "$QUESTION" \
  '{
    messages: [
      {role: "system", content: $system},
      {role: "user", content: ("SOURCES:\n" + $sources + "\n\nQUESTION: " + $question)}
    ],
    max_tokens: 250,
    temperature: 0.3,
    chat_template_kwargs: {enable_thinking: false}
  }' \
| curl -s "$URL" -H "Content-Type: application/json" -d @- \
| jq -r '.choices[0].message.content, "---", "finish_reason: " + .choices[0].finish_reason'
