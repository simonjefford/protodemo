using System;
using System.Net.Http;
using System.Web.UI;
using fourth.people;

namespace aspnetclient
{
    public partial class _Default : Page
    {
        protected void SendToServerBtn_Click(object sender, EventArgs e)
        {
            var b = Employee.CreateBuilder();
            b.Name = this.NameBox.Text;
            b.Age = int.Parse(this.AgeBox.Text);
            b.NiNumber = this.NINumberBox.Text;

            var employee = b.Build();
            var client = new HttpClient();
            var content = new ByteArrayContent(employee.ToByteArray());
            client.PostAsync(new Uri("http://localhost:8080/ReceiveEmployee"), content).Wait();
        }
    }
}