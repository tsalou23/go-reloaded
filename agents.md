# 🤖 AGENTS.md — Go-Reloaded Collaboration Model

## 🎭 Roles

### Operator
- Defines high-level goals and features.  
- Issues tasks to the Agent (via prompts or task files).  
- Prioritizes rule/feature implementation order.  

### Agent (GoReloadedAI)
- Implements tasks using **TDD** (Test-Driven Development).  
- Follows the workflow:
  1. **Analyze** → Understand the problem.  
  2. **Write Tests** → Add Golden/Tricky cases first.  
  3. **Implement Code** → Minimal implementation to pass tests.  
  4. **Refactor & Validate** → Clean code and ensure coverage.  
- Documents progress in `/tasks` with status updates.  

### Auditor
- Reviews test coverage and code quality.  
- Confirms acceptance criteria match the `Analysis.md`.  
- Approves tasks as **Done** when criteria are met.  

---

## 🔄 Workflow

1. **Operator** assigns a task.  
2. **Agent**:
   - Creates or updates a task file in `/tasks` (e.g. `TASK-003-rule-hex-bin.md`).  
   - Writes failing unit tests.  
   - Implements the feature until tests pass.  
   - Marks status as `✅ Done`.  
3. **Auditor** verifies correctness and approves.  

---

## ✅ Acceptance Criteria

- Every rule/feature must:
  - Have **unit tests**.  
  - Pass **Golden** + **Tricky** test cases.  
  - Integrate cleanly into **Pipeline**, **FSM**, and **Hybrid** modes.  

---

## 📂 File Responsibilities

- **Root docs** → Governance (`README.md`, `AGENTS.md`, `CONTRIBUTING.md`, `RELEASE.md`).  
- **`docs/`** → Design docs (`Analysis.md`, `ARCHITECTURE.md`, `TESTING.md`).  
- **`tasks/`** → TDD development logs. One file per milestone.  
- **`tests/`** → Automated coverage (Golden, Tricky, integration).  
- **`internal/`** → Implementation (processors, rules, tokenizer).  

---

## 📜 Example Task Flow

- **Operator**: "Implement hex and bin conversion."  
- **Agent**: Creates `TASK-003-rule-hex-bin.md`, writes failing tests, implements code.  
- **Auditor**: Confirms Golden tests T1/T2 and tricky case C2 pass → approves.  
