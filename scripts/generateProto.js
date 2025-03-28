const { exec } = require('child_process');
const fs = require('fs');
const path = require('path');
const util = require('util');

const execPromise = util.promisify(exec);
const protoDir = path.join(__dirname, '../proto');
const goOutputDir = path.join(__dirname, '../go-server/proto');
const tsOutputDir = path.join(__dirname, '../web-app/src/proto');

// 确保输出目录存在
function ensureDirectoryExists(directory) {
  if (!fs.existsSync(directory)) {
    fs.mkdirSync(directory, { recursive: true });
    console.log(`创建目录: ${directory}`);
  }
}

// 生成Go代码
async function generateGoCode() {
  console.log('开始生成Go代码...');
  ensureDirectoryExists(goOutputDir);
  
  try {
    // 使用本地protoc生成Go代码
    const protocPath = path.join(__dirname, '../protoc/bin/protoc.exe');
    const cmd = `"${protocPath}" --proto_path=${protoDir} `
      + `--go_out=${goOutputDir} `
      + `--go-vtproto_out=${goOutputDir} `
      + `--go-vtproto_opt=features=marshal+unmarshal+size `
      + `--go_opt=paths=source_relative `
      + `--go-vtproto_opt=paths=source_relative `
      + `${protoDir}/*.proto`;
      
    const { stdout, stderr } = await execPromise(cmd);
    if (stdout) console.log(stdout);
    if (stderr) console.error(stderr);
    
    console.log('Go代码生成完成');
  } catch (error) {
    console.error('生成Go代码时出错:', error);
    process.exit(1);
  }
}

// 生成TypeScript代码
async function generateTsCode() {
  console.log('开始生成TypeScript代码...');
  ensureDirectoryExists(tsOutputDir);
  
  try {
    // 获取proto文件的完整路径
    const protoFiles = fs.readdirSync(protoDir)
      .filter(file => file.endsWith('.proto'))
      .map(file => path.join(protoDir, file));
      
    if (protoFiles.length === 0) {
      throw new Error('找不到任何.proto文件');
    }
    
    const protoFilesStr = protoFiles.join(' ');
    console.log('处理的proto文件:', protoFilesStr);
    
    // 第一步：生成静态模块的JavaScript文件
    const pbjs1 = `npx pbjs -t static-module -w commonjs -o "${tsOutputDir}/proto.js" ${protoFilesStr}`;
    console.log(`执行: ${pbjs1}`);
    const res1 = await execPromise(pbjs1);
    if (res1.stderr) console.error(res1.stderr);
    
    // 第二步：从JS生成d.ts类型定义文件
    const pbts = `npx pbts -o "${tsOutputDir}/proto.d.ts" "${tsOutputDir}/proto.js"`;
    console.log(`执行: ${pbts}`);
    const res2 = await execPromise(pbts);
    if (res2.stderr) console.error(res2.stderr);
    
    // 第三步：生成ES模块版本供Vue使用
    const pbjs2 = `npx pbjs -t static-module -w es6 -o "${tsOutputDir}/proto.esm.js" ${protoFilesStr}`;
    console.log(`执行: ${pbjs2}`);
    const res3 = await execPromise(pbjs2);
    if (res3.stderr) console.error(res3.stderr);
    
    // 创建index.ts文件方便导入
    const indexContent = `// TypeScript类型定义
export * from './proto';

// 默认导出ES模块版本
import proto from './proto.esm';
export default proto;
`;
    
    fs.writeFileSync(path.join(tsOutputDir, 'index.ts'), indexContent);
    console.log('创建index.ts文件');
    
    console.log('TypeScript代码生成完成');
  } catch (error) {
    console.error('生成TypeScript代码时出错:', error);
    process.exit(1);
  }
}

// 主函数
async function main() {
  try {
    // 暂时跳过生成Go代码，等Go插件安装好后再启用
    // await generateGoCode();
    console.log('跳过Go代码生成，请先安装Go protobuf插件');
    
    // 生成TypeScript代码
    await generateTsCode();
    
    console.log('所有代码生成完成!');
  } catch (error) {
    console.error('生成过程中出错:', error);
    process.exit(1);
  }
}

// 执行主函数
main(); 