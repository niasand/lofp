## [2026-06-07] 游戏内多处英文未翻译

**现象**：进入游戏后，标题、出口提示、脚本输出等多处显示英文，未进行 i18n 翻译。

**来源**：
1. **UI 字符串**（硬编码在 Go 代码中）：
   - "LEGENDS OF FUTURE PAST" - 标题
   - "The Shattered Realms Await..." - 副标题
   - "Obvious exits" - 出口提示（engine.go, telnet.go, ssh.go）
   - "Type HELP to access the help system" - 帮助提示

2. **脚本字符串**（来自原始 .SCR 文件）：
   - "A trader appears" - 怪物生成文本覆盖（TEXG）
   - "You have been given tunic, breeches and boots" - 来自 FAYDFALL.SCR
   - "A little voice inside your head says..." - 来自 FAYDFALL.SCR

**修复**：
- api.go：标题和欢迎信息使用 i18n.T()
- engine.go：Obvious exits、时间描述、"You see" 等字符串使用 i18n.T()
- telnet.go/ssh.go：Obvious exits 使用 i18n.T()
- scripts.go：doEcho 函数对 ECHO PLAYER 文本使用 i18n.T()
- messages_zh.go：添加所有翻译

**教训**：i18n 覆盖不完整，需要系统性检查所有用户可见字符串。

[AI-REVIEW] Large commit detected: 571 lines added. Consider reviewing for AI Psychosis.
[AI-REVIEW] Large commit detected: 573 lines added. Consider reviewing for AI Psychosis.
[AI-REVIEW] Large commit detected: 201 lines added. Consider reviewing for AI Psychosis.
[AI-REVIEW] Large commit detected: 370 lines added. Consider reviewing for AI Psychosis.
[AI-REVIEW] Large commit detected: 736 lines added. Consider reviewing for AI Psychosis.
[AI-REVIEW] Large commit detected: 738 lines added. Consider reviewing for AI Psychosis.

## [2026-06-07] 创建角色后进入游戏UI"卡住"

**现象**：新建角色后进入游戏，界面看似卡住——没有状态栏（BP/FT/MP/PSI）、没有房间名、没有连接指示器。输入命令后才恢复正常。

**根因**：`api.go` 中 `create_character` 和 `auth_apikey` handler 调用 `EnterRoom` 后直接发送结果，但没有设置 `PlayerState` 和 `PromptIndicators`。前端状态栏渲染条件是 `playerState && (...)`，所以永远不显示。对比 `command` handler 和 telnet 路径都有正确设置。

**修复**：
- `engine/internal/api/api.go`：两处 EnterRoom 结果发送前补上 `result.PlayerState = player` + `result.PromptIndicators = player.PromptIndicators()`
- `frontend/src/components/Terminal.tsx`：补上 `error` 消息类型 handler，避免服务器报错被静默丢弃

**教训**：多个入口路径（WS/telnet/command）处理同一逻辑时，容易遗漏字段赋值。应考虑提取公共的 `enterAndSendResult(session, player)` 函数消除重复。
[AI-REVIEW] Large commit detected: 662 lines added. Consider reviewing for AI Psychosis.
[AI-REVIEW] Large commit detected: 1184 lines added. Consider reviewing for AI Psychosis.
