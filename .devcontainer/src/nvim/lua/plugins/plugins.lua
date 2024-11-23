return {
	{
		"LazyVim/LazyVim",
		opts = {
			colorscheme = "catppuccin",
		},
	},
	{
		"iamcco/markdown-preview.nvim",
		cmd = { "MarkdownPreviewToggle", "MarkdownPreview", "MarkdownPreviewStop" },
		build = function()
			vim.fn["mkdp#util#install"]()
		end,
		keys = {
			{
				"<leader>cp",
				ft = "markdown",
				"<cmd>MarkdownPreviewToggle<cr>",
				desc = "Markdown Preview",
			},
		},
		config = function()
			vim.cmd([[do FileType]])
		end,
	},
	{
		"folke/snacks.nvim",
		opts = {
			terminal = {
				win = {
					position = "top",
					height = 0.25,
				},
			},
		},
		keys = {
			{
				"<C-S-/>",
				function()
					Snacks.terminal.toggle(nil, { win = { position = "float", height = 0.9, width = 0.9 } })
				end,
				desc = "Toggle Terminal",
			},
		},
	},
	{
		"neovim/nvim-lspconfig",
		opts = {
			servers = {
				gopls = {
					settings = {
						gopls = {
							analyses = {
								shadow = false,
							},
							usePlaceholders = false,
						},
					},
				},
			},
		},
	},
	{
		-- NOTE: something broke and char = "|" no longer works
		"lukas-reineke/indent-blankline.nvim",
		opts = function()
			return {
				indent = {},
			}
		end,
	},
}
