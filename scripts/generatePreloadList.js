import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

// 获取当前文件的完整路径
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// 需要扫描的目录（相对于public目录）
const directories = [
    'avatar',
    'background',
    'currentprojects',
    'publications',
    'logo'
]

// 支持的图片格式
const imageExtensions = ['.jpg', '.jpeg', '.png', '.webp', '.gif']

function generateList() {
    const result = []

    // 遍历每个目录
    directories.forEach(dir => {
        const fullPath = path.join(__dirname, '../public', dir)

        if (fs.existsSync(fullPath)) {
            // 递归读取所有文件
            const files = walkSync(fullPath)

            files.forEach(file => {
                const ext = path.extname(file).toLowerCase()
                if (imageExtensions.includes(ext)) {
                    // 生成相对路径（去掉public前缀）
                    const relativePath = path.relative(
                        path.join(__dirname, '../public'),
                        file
                    ).replace(/\\/g, '/')

                    result.push(`/${relativePath}`)
                }
            })
        }
    })

    // 写入到src目录
    fs.writeFileSync(
        path.join(__dirname, '../src/preload-images.json'),
        JSON.stringify(result, null, 2)
    )

    console.log(`成功生成预加载图片列表，共 ${result.length} 张图片`)
}

// 递归遍历目录函数
function walkSync(dir, filelist = []) {
    fs.readdirSync(dir).forEach(file => {
        const dirFile = path.join(dir, file)
        if (fs.statSync(dirFile).isDirectory()) {
            filelist = walkSync(dirFile, filelist)
        } else {
            filelist.push(dirFile)
        }
    })
    return filelist
}

generateList()