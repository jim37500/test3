import axios from 'axios';

export default class CalenderService {
  // 取得行事曆
  static GetCalendar() {
    return axios.get('/api/calendar').then((o) => o.data);
  }

  // 新增行事曆
  static AddCalendar(Name, Date, Todo) {
    return axios.post('/api/calendar', { Name, Date, Todo });
  }

  // 更新行事曆
  static UpdateCalendar(id, Name, Date, Todo) {
    return axios.put(`/api/calendar/${id}`, { Name, Date, Todo });
  }

  // 刪除行事曆
  static DeleteCalendar(id) {
    return axios.delete(`/api/calendar/${id}`);
  }
}
