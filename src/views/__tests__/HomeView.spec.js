// eslint-disable-next-line object-curly-newline
import { describe, it, expect, vi } from 'vitest';

// import { mount, DOMWrapper } from '@vue/test-utils';
import { mount } from '@vue/test-utils';
import HomeView from '../HomeView.vue';
import MyGlobal from '../../__tests__/_Global';
import CalendarService from '../../service/CalendarService';

vi.mock('../../service/CalenderService');
CalendarService.GetCalendar = vi.fn().mockResolvedValue([]);
let wrapper;

describe('HomeView', () => {
  beforeEach(() => {
    CalendarService.AddCalendar = vi.fn().mockResolvedValue({});
    CalendarService.UpdateCalendar = vi.fn().mockResolvedValue({});
    window.Swal = { fire: vi.fn().mockResolvedValue({}) };
    wrapper = null;
  });

  it('清除對話框的資料 並打開對話框', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });
    const newTodo = {
      Name: '',
      Date: '',
      Todo: '',
      ID: 0,
    };

    await wrapper.vm.OpenDialog();

    expect(wrapper.vm.IsShowDialog).toBeTruthy();
    expect(wrapper.vm.NewTodo).toStrictEqual(newTodo);
  });

  it('CheckData 檢查資料是否都填妥 若有 則 回傳成功', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });
    const data = wrapper.vm.NewTodo;
    data.Name = '小明';
    data.Date = '2024/01/24';
    data.Todo = '出差';

    const result = await wrapper.vm.CheckData();

    expect(result).toBeTruthy();
  });

  it('CheckData 檢查資料是否都填妥 若有一個沒有 則 回傳失敗', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });
    const data = wrapper.vm.NewTodo;
    data.Name = '小明';
    data.Date = '2024/01/24';
    data.Todo = '';

    const result = await wrapper.vm.CheckData();

    expect(result).toBeFalsy();
  });

  it('LoadData 載入資料', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });

    await wrapper.vm.LoadData();

    expect(CalendarService.GetCalendar).toHaveBeenCalled();
    expect(wrapper.vm.IsShowDialog).toBeFalsy();
  });

  it('SaveEvent 當ID有值 會更新資料 並呼叫Swal函式', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });

    const data = wrapper.vm.NewTodo;

    data.Name = '小明';
    data.Date = '2024/01/24';
    data.Todo = '出差';
    data.ID = 1;

    await wrapper.vm.SaveEvent();

    expect(CalendarService.UpdateCalendar).toHaveBeenCalled();
    expect(window.Swal.fire).toHaveBeenCalledWith('更新成功!');
    expect(wrapper.vm.IsShowDialog).toBeFalsy();
  });

  it('SaveEvent 當資料未填妥 則回傳null 並呼叫Swal函式', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });
    const data = wrapper.vm.NewTodo;
    data.Name = '小明';
    data.Date = '2024/01/24';
    data.Todo = '';

    await wrapper.vm.SaveEvent();

    expect(window.Swal.fire).toHaveBeenCalledWith('請確實填妥資料', '', 'error');
    expect(wrapper.vm.IsShowDialog).toBeFalsy();
  });

  it('SaveEvent 當ID沒有值 會新增資料 並呼叫Swal函式', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });
    const data = wrapper.vm.NewTodo;
    data.Name = '小明';
    data.Date = '2024/01/24';
    data.Todo = '放假';

    await wrapper.vm.SaveEvent();

    expect(CalendarService.AddCalendar).toHaveBeenCalled();
    expect(window.Swal.fire).toHaveBeenCalledWith('新增成功!');
    expect(wrapper.vm.IsShowDialog).toBeFalsy();
  });

  it('DeleteEvent 跳出選項 詢問是否刪除資料', async () => {
    wrapper = mount(HomeView, { global: MyGlobal });

    await wrapper.vm.DeleteEvent();

    expect(window.Swal.fire).toHaveBeenCalledWith({
      title: '確定刪除此行程?',
      showDenyButton: false,
      showCancelButton: true,
      confirmButtonText: '刪除',
      cancelButtonText: '取消',
    });
  });
});
