local map = vim.keymap.set

map("n", "gct", "O<esc>VcTODO: <esc><cmd>normal gcc<cr>A", { desc = "Add TODO comment above current line" })
-- HACK: is there a better way to do this?
map("n", "gcat", "<cmd>normal gct<cr><esc>0v$d$a <esc>p$jddA", { desc = "Add TODO comment at the end of line" })
map("n", "<leader>n", "<cmd>Neotree focus<cr>", { desc = "focus Neotree" })
