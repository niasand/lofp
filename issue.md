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
