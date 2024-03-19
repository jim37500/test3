// eslint-disable-next-line object-curly-newline
import { describe, it, expect, vi } from 'vitest';
import axios from 'axios';
import CalendarService from '../CalendarService';

const fakeData = { Name: 'fakeName', Date: 'fakeDate', Todo: 'fakeTodo' };

describe('CalendarService', () => {
  it('讀取資料', async () => {
    axios.get = vi.fn().mockResolvedValue({ data: 'fakeData' });

    const result = await CalendarService.GetCalendar();

    expect(result).toBe('fakeData');
    expect(axios.get).toHaveBeenCalledWith('/api/calendar');
  });

  it('新增資料', async () => {
    axios.post = vi.fn().mockResolvedValue();

    await CalendarService.AddCalendar('fakeName', 'fakeDate', 'fakeTodo');

    expect(axios.post).toHaveBeenCalledWith('/api/calendar', fakeData);
  });

  it('更新資料', async () => {
    axios.put = vi.fn().mockResolvedValue();

    await CalendarService.UpdateCalendar(1, 'fakeName', 'fakeDate', 'fakeTodo');

    expect(axios.put).toHaveBeenCalledWith('/api/calendar/1', fakeData);
  });

  it('刪除資料', async () => {
    axios.delete = vi.fn().mockResolvedValue();

    await CalendarService.DeleteCalendar(1);

    expect(axios.delete).toHaveBeenCalledWith('/api/calendar/1');
  });
});
