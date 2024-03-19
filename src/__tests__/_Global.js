/* c8 ignore start */
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import { defaultOptions } from 'primevue/config';

export default {
  components: {
    Button,
    Dialog,
    InputText,
    FullCalendar: 'FullCalendar',
  },
  mocks: {
    $primevue: {
      config: defaultOptions,
    },
  },
};
/* c8 ignore stop */
