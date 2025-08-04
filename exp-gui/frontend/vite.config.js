import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path,{ resolve } from 'path';
// https://vite.dev/config/
export default defineConfig({
    plugins: [vue()],
    port: 8080,
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src'), // 使用 @ 指向 src 目录
            '/@': path.resolve(__dirname, './'), // 使用 @ 指向 frontend 目录
        },
    },
    css: {//arco design时用的
        preprocessorOptions: {
            less: {
                modifyVars: {
                    'arcoblue-6': "#165DFF",//
                },
                javascriptEnabled: true,
            },
        },
    },
})