/* eslint-disable import/no-extraneous-dependencies */
import './assets/main.css';
import 'primeflex/primeflex.css';
import Swal from 'sweetalert2';

import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/lara-light-green/theme.css';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
// import '@fullcalendar/core/vdom';
import FullCalendar from '@fullcalendar/vue3';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import listPlugin from '@fullcalendar/list';
import interactionPlugin from '@fullcalendar/interaction';
import App from './App.vue';
import router from './router';

const app = createApp(App);
window.dayGridPlugin = dayGridPlugin;
window.timeGridPlugin = timeGridPlugin;
window.listPlugin = listPlugin;
window.interactionPlugin = interactionPlugin;
window.Swal = Swal;

app.use(router);
app.use(PrimeVue);
app.component('Button', Button);
app.component('Dialog', Dialog);
app.component('InputText', InputText);
app.component('FullCalendar', FullCalendar);

app.mount('#app');
