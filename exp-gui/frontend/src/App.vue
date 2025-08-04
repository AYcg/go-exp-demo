<template>
  <div class="app-container">
    <!-- 顶部按钮行 -->
    <div class="top-row">
      <a-trigger trigger="click" :unmount-on-close="false">
        <a-button class="action-btn">设置</a-button>
        <template #content>
          <div class="setting">
            <a-button @click="showProxyModal" class="action-btn">代理设置</a-button>
            <a-button @click="showHeaderModal" class="action-btn">请求头设置</a-button>
          </div>
        </template>
      </a-trigger>
      <a-button @click="handleClick">关于</a-button>
      <!-- 请求头设置弹窗 -->
      <a-modal
          v-model:visible="headerVisible"
          title="请求头设置"
          :mask-closable="false"
          width="600px"
      >
        <a-table
            :data="headers"
            :pagination="false"
            :bordered="true"
            row-key="id"
        >
        <template #columns>
          <a-table-column title="Header Key">
            <template #cell="{ record }">
              <a-input
                  v-model="record.key"
                  placeholder="请输入 Header Key"
                  @change="validateHeader(record)"
              />
            </template>
          </a-table-column>

          <a-table-column title="Header Value">
            <template #cell="{ record }">
              <a-input
                  v-model="record.value"
                  placeholder="请输入 Header Value"
                  @change="validateHeader(record)"
              />
            </template>
          </a-table-column>

          <a-table-column title="操作" :width="120">
            <template #cell="{ rowIndex }">
              <a-button
                  type="text"
                  status="danger"
                  @click="removeHeader(rowIndex)"
              >
                <template #icon><icon-delete /></template>
              </a-button>
            </template>
          </a-table-column>
        </template>
        </a-table>

        <div class="add-header-btn">
          <a-button @click="addHeader" type="outline">
            <template #icon><icon-plus /></template>添加请求头
          </a-button>
        </div>

        <template #footer>
          <a-button @click="handleHeaderCancel">取消</a-button>
          <a-button type="primary" @click="handleHeaderOk">确定</a-button>
        </template>
      </a-modal>



      <!-- 代理设置弹窗 -->
      <a-modal
          v-model:visible="proxyVisible"
          title="代理设置"
          :mask-closable="false"
          width="600px"
      >
        <a-form :model="proxyForm">
          <a-form-item label="代理地址">
            <a-input
                v-model="proxyForm.address"
                placeholder="例如: http://127.0.0.1:8080"
            />
            <a-button
                type="outline"
                @click="testProxy"
                :disabled="!proxyForm.address.trim()"
            >
              测试代理
            </a-button>
          </a-form-item>
        </a-form>

        <template #footer>
          <a-button @click="handleProxyCancel">取消</a-button>
          <a-button type="primary" @click="handleProxyOk">确定</a-button>
        </template>
      </a-modal>

<!--      关于说明弹窗-->
      <a-modal v-model:visible="visible" @ok="handleOk" @cancel="handleCancel">
        <template #title>
          关于 长歌漏洞利用工具1
        </template>

          <div class="about-section">
            <h3><icon-question-circle-fill /> 使用须知</h3>
            <ol>
              <li>本工具仅限授权安全测试使用</li>
              <li>使用者须遵守《网络安全法》及相关法律</li>
              <li>禁止用于未授权测试和非法用途</li>
              <li>商业使用需获得作者授权</li>
            </ol>
          </div>

          <a-divider/>

          <div class="about-section author-info">
            <h3><icon-user/> 作者信息</h3>
            <div class="info-grid">
              <span>开发者：</span><span>changge</span>
              <br><span>版本：</span><span>v1.0.0</span>
              <br><span>联系方式：</span><span>changge1001@gmail.com</span>
            </div>
          </div>


        <template #footer>
          <a-button type="primary" @click="handleOk">我已阅读并同意</a-button>
        </template>
      </a-modal>

    </div>

    <!-- 主控制区 -->
    <div class="control-row">
      <input
          id="url"
          v-model="targetUrl"
          class="url-input"
          type="text"

      >

      <select v-model="selectedPoc" class="vuln-select">
        <option v-for="poc in pocList" :key="poc.value" :value="poc.value">
          {{ poc.label }}
        </option>
      </select>

      <button @click="poc" class="action-btn">漏洞检测</button>
      <button @click="info"  class="action-btn">漏洞描述</button>
      <button @click="clearActiveTab" class="action-btn">清空面板</button>
    </div>
  </div>
<!--  子区域-->
  <div class="tab-area">
    <a-tabs v-model:active-key="activeTabKey">
      <a-tab-pane key="1">
        <template #title>
          <icon-calendar/> 基本信息
        </template>
        <div class="info-tab">
          <div class="tab-content" v-html="tabContents['1']"></div>
        </div>
      </a-tab-pane>
      <a-tab-pane key="2">
        <template #title>
          <icon-clock-circle/> 命令执行
        </template>
        <rce :target-url="targetUrl" :select-poc="selectedPoc" @execute-cmd="handleRceResult" />
        <div class="rce-tab">
          <div class="tab-content" v-html="tabContents['2']"></div>
        </div>
      </a-tab-pane>
      <a-tab-pane key="3">
        <template #title>
          <icon-user/> 文件上传
        </template>
        <upload :target-url="targetUrl" :select-poc="selectedPoc" @file-upload="handleUploadResult" />
        <div class="rce-tab">
          <div class="tab-content" v-html="tabContents['3']"></div>
        </div>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup>

import rce from '@/components/Rce.vue'
import upload from '@/components/Upload.vue'
import { Message } from '@arco-design/web-vue';
import {ToolService} from "../bindings/changeme/service";
import {Service} from "../bindings/changeme/service/tools";
import { ref, onMounted } from 'vue'
import { IconDelete, IconPlus } from '@arco-design/web-vue/es/icon'



const targetUrl = ref('')
// 漏洞类型列表
const pocList = ref([])

// 获取漏洞列表
const loadPocList = async () => {
  try {
    pocList.value = await ToolService.GetPocList()
  } catch (error) {
    Message.error('加载漏洞列表失败: ' + error.message)
    // 备用方案：使用默认列表
    pocList.value = ['all', 'demo1', 'demo2', 'demo3']
  }
}

// 基本信息
const loadInfo = ref('')
const info = async () => {
  try {
    tabContents.value[1] = await ToolService.GetInfo()
  } catch (error) {
    Message.error('获取基本信息失败: ' + error.message)
  }
}

// tab基本信息
const activeTabKey = ref('1');
const tabContents = ref({
  '1': '',  // 基本信息
  '2': '命令执行结果将显示在这里',
  '3': '文件上传区域'
});



// 关于信息弹窗
const visible = ref(false);
const handleClick = () => {
  visible.value = true;
};
const handleOk = () => {
  visible.value = false;
};
const handleCancel = () => {
  visible.value = false;
}


// poc检测
const selectedPoc = ref('all');
const poc = () => {
  var action = 'check';
  let url = document.getElementById('url').value;
  if (!url) {
    Message.error({
      content: '请输入有效的URL地址',
      duration: 2000,
      closable: true
    });
    return;
  }

  ToolService.POC(url, selectedPoc.value, action,"","").then((result) => {
    tabContents.value[activeTabKey.value]  = result;
  });

}


// 清空当前Tab
const clearActiveTab = () => {
  tabContents.value = {
    ...tabContents.value,
    [activeTabKey.value]: `<div class="empty-tip">${getTabName(activeTabKey.value)}内容已清空</div>`
  };
  Message.success(`已清空${getTabName(activeTabKey.value)}面板`);
};

// 获取Tab名称
const getTabName = (key) => {
  const names = { '1': '基本信息', '2': '命令执行', '3': '文件上传' };
  return names[key] || '';
};



// 设置请求头
// 设置请求头相关
const headerVisible = ref(false);
const headers = ref([
  { id: Date.now(), key: '', value: '' } // 加 id
]); // 默认一个空行

const showHeaderModal = () => {

  if (headers.value.length === 0) {
    headers.value.push({id: Date.now() + Math.random(), key: '', value: '' });
  }

  headerVisible.value = true;
};

const addHeader = () => {
  headers.value.push({
    id: Date.now() + Math.random(), // 防止重复
    key: '',
    value: ''
  });
};


const removeHeader = (index) => {
  headers.value.splice(index, 1);
};

const validateHeader = (header) => {
  // 验证逻辑...
};

const handleHeaderOk = async () => {
  try {
    // 过滤并转换格式
    const headersObj = {};
    alert(headers.value)
    headers.value.forEach(item => {
      if (item.key?.trim() && item.value?.trim()) {
        headersObj[item.key.trim()] = item.value.trim();
      }
    });

    alert("最终 headersObj:", headersObj);


    // 调用后端保存（注意方法名要与后端一致）
    await Service.SetGlobalHeaders(headersObj);


    Message.success(`已设置 ${Object.keys(headersObj).length} 个请求头`);
    headerVisible.value = false;
  } catch (error) {
    Message.error(error.message);
  }
};

// 初始化加载已有请求头和基本信息
onMounted(async () => {
  info();
  await loadPocList();
  try {
    const savedHeaders = await Service.GetGlobalHeaders();
    if (savedHeaders && Object.keys(savedHeaders).length > 0) {
      headers.value = Object.entries(savedHeaders).map(([key, value]) => ({
        id: Date.now() + Math.random(),
        key,
        value
      }));
    }
  } catch (error) {
    console.error('加载请求头失败:', error);
  }
});


const handleHeaderCancel = () => {
  headerVisible.value = false;
};




// 代理设置相关
const proxyVisible = ref(false);
const proxyForm = ref({
  address: ''
});

// 打开代理设置弹窗
const showProxyModal = async () => {
  try {
    // 获取当前代理设置
    const currentProxy = await Service.GetProxy();
    proxyForm.value.address = currentProxy || '';
    proxyVisible.value = true;
  } catch (error) {
    Message.error('获取代理设置失败: ' + error.message);
  }
};

// 保存代理设置
const handleProxyOk = async () => {
  try {
    const proxyAddress = proxyForm.value.address.trim();

    // 简单验证代理格式
    if (proxyAddress && !/^https?:\/\/.+\:\d+$/.test(proxyAddress)) {
      throw new Error('请输入有效的代理地址，如: http://127.0.0.1:8080');
    }

    await Service.SetProxy(proxyAddress);
    Message.success(proxyAddress ? '代理设置成功' : '已清除代理设置');
    proxyVisible.value = false;
  } catch (error) {
    Message.error('保存失败: ' + error.message);
  }
};

// 取消代理设置
const handleProxyCancel = () => {
  proxyVisible.value = false;
};

// 可选：代理测试功能
const testProxy = async () => {
  try {
    const proxyAddress = proxyForm.value.address.trim();
    if (!proxyAddress) {
      throw new Error('请先输入代理地址');
    }

    const [isValid, message] = await Service.TestProxy(proxyAddress);
    if (isValid) {
      Message.success('代理测试成功: ' + message);
    } else {
      Message.error('代理测试失败: ' + message);
    }
  } catch (error) {
    Message.error('测试失败: ' + error.message);
  }
};


// 处理子控件rce模块传入的数据
const handleRceResult = (result) => {
  tabContents.value['2'] = result
}
// 处理子控件文件上传模块传入的数据
const handleUploadResult = (result) => {
  tabContents.value['3'] = result
}

</script>

<style>
/* 主容器 */
.app-container {
  padding: 10px;
  width: 100%;
  max-width: 760px;
  margin: 0 auto;
}

/* 顶部按钮行 */
.top-row {
  display: flex;
  gap: 10px;
  margin-bottom: 5px;
}

/* 主控制行 */
.control-row {
  display: flex;
  gap: 8px;
  width: 100%;

}

/* 统一输入控件样式 */
.url-input,
.vuln-select,
.action-btn {
  padding: 6px 12px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  font-size: 14px;
}
.btn{
  margin-top: 1px;
}

/* URL输入框 */
.url-input {
  flex: 2;
  min-width: 180px;
}

/* 漏洞选择框 */
.vuln-select {
  flex: 1;
  min-width: 100px;
}

/* 操作按钮 */
.action-btn {
  background: #f5f5f5;
  cursor: pointer;
  white-space: nowrap;
}

/* 设置列表 */
.setting{
  display: flex;
  flex-direction: column;
  width: 100px;

}

.proxy-test-btn {
  margin-top: 12px;
  text-align: center;
}
.add-header-btn {
  margin-top: 16px;
  text-align: center;
}

/* 响应式调整 */
@media (max-width: 600px) {
  .control-row {
    flex-wrap: wrap;
  }

  .url-input,
  .vuln-select,
  .action-btn {
    width: 80%;
  }
}

// table 样式
.tab-area{
  width: 780px;

  border: 1px solid #d9d9d9;
}
.tab-content{

}
.rce-tab{
  width: 700px;
  height: 300px;
  background: #ffffff;
  border: 1px solid #d9d9d9;
  padding: 20px 20px;
  margin: 10px 18px;
}
.info-tab{
  width: 700px;
  height: 300px;
  padding: 20px 20px;
  margin: 2px 18px;
}


</style>