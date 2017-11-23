# cloudgo-io

贡献度分工比：  
15331149李欢： 50%   
15331147李海宏： 50%  


### 基本要求：

实验截图：

![index](pictures/4.png)

1.

   ![静态文件服务](pictures/main.css.png)

   ![使用/static/](pictures/static.png)

2. js 访问

   利用js获取系统时间，每1秒更新一次：

   **clock.js**:

   ```javascript
   $(document).ready(function() {
       setInterval(function() {
         var dateNtime = new Date();
         var year = dateNtime.getFullYear();
         var month = dateNtime.getMonth() + 1;
         var day = dateNtime.getDate();
         var hours = dateNtime.getHours();
         var minute = dateNtime.getMinutes();
         var second = dateNtime.getSeconds();

         var dateStr = month + "/" + day + ", " + year;
         var timeStr = hours + ":" + minute + ":" + second;
         $('.index-date').html("Today is " + dateStr);
         $('.index-time').html("The time is " + timeStr);
       }, 1000)
   })
   ```

3. 提交注册单，还有表格

   在"/register"页面下进行注册，提交表单：

   ![填写表单](pictures/register.png)

   ![表格](pictures/table.png)  

4. 对 `/unknown` 给出开发中的提示，返回码 `5xx`

   ![unknown](pictures/unknown.png)

   ​
