// 忽略 PrimeVue CSS錯誤
const originalConsoleError = console.error;
console.error = (...data) => {
  if (typeof data[0]?.toString === 'function' && data[0].toString().includes('Error: Could not parse CSS stylesheet')) return;
  originalConsoleError(...data);
};
