local map = vim.keymap.set

map("n", "gct", "O<esc>VcTODO: <esc><cmd>normal gcc<cr>A", { desc = "Add TODO comment above current line" })
map("n", "gcT", "o<esc>VcTODO: <esc><cmd>normal gcc<cr>A", { desc = "Add TODO comment below current line" })
map("n", "<leader>n", "<cmd>Neotree focus<cr>", { desc = "focus Neotree" })
