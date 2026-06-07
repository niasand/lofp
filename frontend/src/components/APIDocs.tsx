import { t } from '../i18n'

export default function APIDocs({ onBack }: { onBack: () => void }) {
  return (
    <div className="flex items-start justify-center h-full p-8 overflow-y-auto">
      <div className="max-w-3xl w-full font-mono">
        <div className="flex items-center justify-between mb-6">
          <h1 className="text-amber-500 text-2xl font-bold">{t("api.title")}</h1>
          <button onClick={onBack} className="text-gray-400 hover:text-white text-sm">{t("api.back")}</button>
        </div>

        <div className="bg-[#111] border border-[#333] rounded-lg p-4 mb-6">
          <p className="text-gray-300 text-sm leading-relaxed">
            {t("api.intro")}
          </p>
          <p className="text-gray-400 text-xs mt-3">
            <a href="/bot-agent-spec.md" className="text-amber-400 hover:text-amber-300 underline">
              {t("api.agentSpec")}
            </a>
            &nbsp;{t("api.agentSpecSuffix")}
          </p>
        </div>

        <div className="space-y-6 text-sm">

          <section>
            <h2 className="text-amber-400 text-lg font-bold mb-2">{t("api.getKeyTitle")}</h2>
            <ol className="text-gray-300 space-y-2 ml-4 list-decimal">
              <li>{t("api.getKeyStep1a")} <a href="https://lofp.metavert.io" className="text-amber-400 hover:underline">lofp.metavert.io</a></li>
              <li>{t("api.getKeyStep2a")} <span className="text-amber-300">&#9881; Bot</span> {t("api.getKeyStep2b")}</li>
              <li>{t("api.getKeyStep3a")} <span className="text-amber-300">{t("api.generateKey")}</span> {t("api.getKeyStep3b")}</li>
              <li><strong>{t("api.copyKeyHint")}</strong> {t("api.getKeyStep4b")}</li>
              <li>{t("api.regenerateHint")}</li>
            </ol>
          </section>

          <section>
            <h2 className="text-amber-400 text-lg font-bold mb-2">{t("api.connTitle")}</h2>
            <ol className="text-gray-300 space-y-2 ml-4 list-decimal">
              <li>{t("api.connStep1")} <code className="text-green-400 bg-[#0a0a0a] px-1 rounded">wss://lofp.metavert.io/ws/game</code></li>
              <li>{t("api.welcomeMessage")}</li>
              <li>{t("api.connStep3a")} <code className="text-green-400 bg-[#0a0a0a] px-1 rounded">auth_apikey</code> {t("api.connStep3b")}</li>
              <li>{t("api.loginSuccess")}</li>
              <li>{t("api.connStep5a")} <code className="text-green-400 bg-[#0a0a0a] px-1 rounded">command</code> {t("api.connStep5b")}</li>
            </ol>
          </section>

          <section>
            <h2 className="text-amber-400 text-lg font-bold mb-2">{t("api.msgFormatTitle")}</h2>
            <p className="text-gray-400 mb-3">{t("api.msgFormatDesc1")} <code className="text-green-400">type</code> {t("api.msgFormatDesc2")} <code className="text-green-400">data</code> {t("api.msgFormatDesc3")}</p>

            <h3 className="text-green-400 font-bold mb-1">{t("api.sending")}</h3>
            <pre className="bg-[#0a0a0a] border border-[#333] rounded p-3 text-xs text-gray-300 overflow-x-auto mb-4">{`// Authenticate
{"type": "auth_apikey", "data": {"key": "lofp_abc123..."}}

// Send a game command
{"type": "command", "data": {"input": "look"}}
{"type": "command", "data": {"input": "attack skeleton"}}
{"type": "command", "data": {"input": "say Hello everyone!"}}`}</pre>

            <h3 className="text-green-400 font-bold mb-1">{t("api.receiving")}</h3>
            <pre className="bg-[#0a0a0a] border border-[#333] rounded p-3 text-xs text-gray-300 overflow-x-auto mb-4">{`// Auth result
{"type": "auth_result", "data": {"success": true, "character": "MyBot"}}

// Game output (from your commands)
{"type": "result", "data": {"messages": ["[City Gate]", "You stand at..."], "roomName": "[City Gate]"}}

// Broadcast (from other players/monsters)
{"type": "broadcast", "data": {"messages": ["A skeleton attacks you!"]}}`}</pre>

            <h3 className="text-green-400 font-bold mb-1">{t("api.responseFields")}</h3>
            <div className="overflow-x-auto">
              <table className="w-full text-xs">
                <thead><tr className="text-gray-400 border-b border-[#333]">
                  <th className="text-left py-1 pr-3">{t("api.field")}</th><th className="text-left py-1">{t("api.description")}</th>
                </tr></thead>
                <tbody className="text-gray-300">
                  <tr className="border-b border-[#222]"><td className="py-1 pr-3 text-amber-300">messages</td><td>{t("api.linesArray")}</td></tr>
                  <tr className="border-b border-[#222]"><td className="py-1 pr-3 text-amber-300">roomName</td><td>{t("api.roomName")}</td></tr>
                  <tr className="border-b border-[#222]"><td className="py-1 pr-3 text-amber-300">roomDesc</td><td>{t("api.roomDesc")}</td></tr>
                  <tr className="border-b border-[#222]"><td className="py-1 pr-3 text-amber-300">exits</td><td>{t("api.exitsArray")}</td></tr>
                  <tr className="border-b border-[#222]"><td className="py-1 pr-3 text-amber-300">items</td><td>{t("api.itemsArray")}</td></tr>
                  <tr className="border-b border-[#222]"><td className="py-1 pr-3 text-amber-300">playerState</td><td>{t("api.stats")}</td></tr>
                  <tr><td className="py-1 pr-3 text-amber-300">promptIndicators</td><td>{t("api.statusCodes")}</td></tr>
                </tbody>
              </table>
            </div>
          </section>

          <section>
            <h2 className="text-amber-400 text-lg font-bold mb-2">{t("api.botBehaviorTitle")}</h2>
            <ul className="text-gray-300 space-y-1 ml-4 list-disc">
              <li>{t("api.botBehavior1a")} <span className="text-amber-300">[Bot]</span> {t("api.botBehavior1b")}</li>
              <li>{t("api.rateLimit")}</li>
              <li>{t("api.chatFlood")}</li>
              <li>{t("api.gmCommandsHint")}</li>
              <li>{t("api.normalRules")}</li>
            </ul>
          </section>

          <section>
            <h2 className="text-amber-400 text-lg font-bold mb-2">{t("api.codeExamplesTitle")}</h2>

            <h3 className="text-green-400 font-bold mb-1 mt-4">Python</h3>
            <pre className="bg-[#0a0a0a] border border-[#333] rounded p-3 text-xs text-gray-300 overflow-x-auto mb-4">{`# pip install websocket-client
import json, os, websocket

API_KEY = os.environ["LOFP_API_KEY"]
ws = websocket.WebSocket()
ws.connect("wss://lofp.metavert.io/ws/game")

# Skip welcome message
ws.recv()

# Authenticate
ws.send(json.dumps({"type": "auth_apikey", "data": {"key": API_KEY}}))
result = json.loads(ws.recv())
print("Auth:", result)

# Send a command
ws.send(json.dumps({"type": "command", "data": {"input": "look"}}))
result = json.loads(ws.recv())
for line in result.get("data", {}).get("messages", []):
    print(line)`}</pre>

            <h3 className="text-green-400 font-bold mb-1">Node.js</h3>
            <pre className="bg-[#0a0a0a] border border-[#333] rounded p-3 text-xs text-gray-300 overflow-x-auto mb-4">{`// npm install ws
const WebSocket = require("ws");
const ws = new WebSocket("wss://lofp.metavert.io/ws/game");

ws.on("open", () => {
  ws.send(JSON.stringify({
    type: "auth_apikey",
    data: { key: process.env.LOFP_API_KEY }
  }));
});

ws.on("message", (raw) => {
  const msg = JSON.parse(raw);
  if (msg.type === "auth_result" && msg.data.success) {
    ws.send(JSON.stringify({
      type: "command", data: { input: "look" }
    }));
  }
  if (msg.data?.messages) {
    msg.data.messages.forEach(line => console.log(line));
  }
});`}</pre>

            <h3 className="text-green-400 font-bold mb-1">TypeScript</h3>
            <pre className="bg-[#0a0a0a] border border-[#333] rounded p-3 text-xs text-gray-300 overflow-x-auto mb-4">{`// npm install ws @types/ws tsx
import WebSocket from "ws";

interface GameMessage {
  type: string;
  data: {
    success?: boolean;
    character?: string;
    messages?: string[];
    roomName?: string;
    playerState?: { bodyPoints: number; maxBodyPoints: number };
  };
}

const ws = new WebSocket("wss://lofp.metavert.io/ws/game");

ws.on("open", () => {
  ws.send(JSON.stringify({
    type: "auth_apikey",
    data: { key: process.env.LOFP_API_KEY }
  }));
});

ws.on("message", (raw: WebSocket.Data) => {
  const msg: GameMessage = JSON.parse(raw.toString());
  if (msg.type === "auth_result" && msg.data.success) {
    console.log(\`Logged in as \${msg.data.character}\`);
    ws.send(JSON.stringify({
      type: "command", data: { input: "look" }
    }));
  }
  for (const line of msg.data.messages ?? []) {
    console.log(line);
  }
});`}</pre>
          </section>

          <section>
            <h2 className="text-amber-400 text-lg font-bold mb-2">{t("api.fullExamplesTitle")}</h2>
            <p className="text-gray-300 text-sm">
              {t("api.fullExamplesDesc1")}{' '}
              <a href="https://github.com/jonradoff/lofp/tree/main/bots" className="text-amber-400 hover:underline">
                {t("api.fullExamplesLink")}
              </a>{t("api.fullExamplesDesc2")}
            </p>
          </section>

        </div>

        <div className="mt-8 pt-4 border-t border-[#333] text-gray-600 text-xs text-center">
          {t("api.footer")}
        </div>
      </div>
    </div>
  )
}
