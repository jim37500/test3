/* c8 ignore start */

<template>
  <div>
    <FullCalendar :options="CalendarOptions" class="" />
    <!-- 新增行事曆按鈕 -->
    <div class="card block mt-3 justify-content-center">
      <Button data-test="button_open-dialog" label="新增行事曆" icon="pi pi-user" @click="OpenDialog" />
    </div>
  </div>

  <!-- 行事曆的對話框 -->
  <Dialog data-test="dialog" v-model:visible="IsShowDialog" modal :pt="{ mask: { style: 'backdrop-filter: blur(2px)' } }">
    <template #container="{ closeCallback }">
      <div class="flex flex-column px-8 py-5 gap-4" style="border-radius: 12px; background-image: radial-gradient(circle at left top, var(--gray-600), var(--gray-700))">
        <div class="inline-flex flex-column gap-2">
          <label for="username" class="text-primary-50 font-semibold">使用者</label>
          <InputText id="username" class="bg-white-alpha-20 border-none p-3 text-primary-50" v-model="NewTodo.Name" />
        </div>
        <div class="inline-flex flex-column gap-2">
          <label for="date" class="text-primary-50 font-semibold">日期</label>
          <InputText id="date" class="bg-white-alpha-20 border-none p-3 text-primary-50" type="date" v-model="NewTodo.Date" />
        </div>
        <div class="inline-flex flex-column gap-2">
          <label for="todo" class="text-primary-50 font-semibold">行程</label>
          <InputText id="todo" class="bg-white-alpha-20 border-none p-3 text-primary-50" v-model="NewTodo.Todo" />
        </div>

        <div class="flex align-items-center gap-2">
          <Button label="儲存" @click="SaveEvent" text class="p-3 w-full text-primary-50 border-1 border-white-alpha-30 hover:bg-white-alpha-10" />

          <Button v-if="NewTodo.ID" label="刪除" @click="DeleteEvent" text class="p-3 w-full text-primary-50 border-1 border-white-alpha-30 hover:bg-white-alpha-10" />
          <Button label="取消" @click="closeCallback" text class="p-3 w-full text-primary-50 border-1 border-white-alpha-30 hover:bg-white-alpha-10" />
        </div>
      </div>
    </template>
  </Dialog>
</template>
/* c8 ignore end */

<script setup>
import { ref, reactive } from 'vue';
import CalendarService from '../service/CalendarService';

// eslint-disable-next-line object-curly-newline
const { dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin, Swal } = window;
const IsShowDialog = ref(false); // 是否顯示行事曆的對話框

// 對話框行事曆的資料格式
const NewTodo = reactive({
  Name: '',
  Date: '',
  Todo: '',
  ID: 0,
});

// 行事曆的資料
const events = [];

// 清空對話框的資料 並 關閉對話框
const OpenDialog = () => {
  NewTodo.Name = '';
  NewTodo.Date = '';
  NewTodo.Todo = '';
  NewTodo.ID = 0;

  IsShowDialog.value = true;
};

// 行事曆設定
const CalendarOptions = reactive({
  plugins: [dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin],
  initialView: 'dayGridMonth',
  height: 800,
  headerToolbar: {
    start: 'prev,next today',
    center: 'title',
    end: 'dayGridMonth,dayGridWeek,listDay',
  },
  dateClick: (info) => {
    OpenDialog();
    NewTodo.Date = info.dateStr;
  },
  selectable: true,
  events,
  eventClick: (info) => {
    info.jsEvent.preventDefault();
    OpenDialog();
    NewTodo.Name = info.event.extendedProps.Name;
    NewTodo.Date = info.event.startStr;
    NewTodo.Todo = info.event.extendedProps.Todo;
    NewTodo.ID = info.event.extendedProps.ID;
  },
});

// 從後端載入資料
const LoadData = () => {
  IsShowDialog.value = false;

  CalendarService.GetCalendar().then((o) => {
    CalendarOptions.events = o;
    CalendarOptions.events.forEach((el) => {
      const item = el;
      item.title = `${el.Name}  ${el.Todo}`;
      item.start = el.Date;
    });
  });
};
LoadData();

// 檢查資料
const CheckData = () => {
  // 若 新增對話框的 Name Date Todo 沒值 回傳否
  if (!NewTodo.Name || !NewTodo.Date || !NewTodo.Todo) {
    return false;
  }

  return true;
};

// 存取資料
const SaveEvent = () => {
  // 若資料未填妥 則 回傳空值
  if (!CheckData()) {
    Swal.fire('請確實填妥資料', '', 'error');
    IsShowDialog.value = false;
    return null;
  }

  // 若 傳入的資料的ID有值 則 更新資料
  if (NewTodo.ID) {
    return CalendarService.UpdateCalendar(NewTodo.ID, NewTodo.Name, NewTodo.Date, NewTodo.Todo)
      .then(LoadData)
      .then(() => Swal.fire('更新成功!'))
      .catch(() => Swal.fire('更新失敗!'));
  }

  // 如果傳入的資料得ID沒值 則 新增資料
  return CalendarService.AddCalendar(NewTodo.Name, NewTodo.Date, NewTodo.Todo)
    .then(LoadData)
    .then(() => Swal.fire('新增成功!'))
    .catch(() => {
      IsShowDialog.value = false;
      Swal.fire('新增失敗!');
    });
};

// 刪除資料
const DeleteEvent = () => {
  IsShowDialog.value = false;
  Swal.fire({
    title: '確定刪除此行程?',
    showDenyButton: false,
    showCancelButton: true,
    confirmButtonText: '刪除',
    cancelButtonText: '取消',
  }).then((result) => {
    // 若 點選"刪除"按鈕 則 刪除資料
    if (result.isConfirmed) {
      CalendarService.DeleteCalendar(NewTodo.ID)
        .then(LoadData)
        .then(() => Swal.fire('已刪除行程!', '', 'success'))
        .catch(() => Swal.fire('刪除失敗!', '', 'error'));
    }
  });
};
</script>
