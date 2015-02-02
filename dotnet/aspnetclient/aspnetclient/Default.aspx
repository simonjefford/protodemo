<%@ Page Title="Home Page" Language="C#" MasterPageFile="~/Site.Master" AutoEventWireup="true" CodeBehind="Default.aspx.cs" Inherits="aspnetclient._Default" %>

<asp:Content runat="server" ID="FeaturedContent" ContentPlaceHolderID="FeaturedContent">
    <section class="featured">
        <div class="content-wrapper">
            <hgroup class="title">
                <h1><%: Title %>.</h1>
                <h2>Protocol buffers demo</h2>
            </hgroup>
        </div>
    </section>
</asp:Content>
<asp:Content runat="server" ID="BodyContent" ContentPlaceHolderID="MainContent">
    <h3>Enter employee:</h3>
    <p>
        Name: <asp:TextBox runat="server" ID="NameBox"></asp:TextBox>
    </p>
    <p>
        Age: <asp:TextBox runat="server" ID="AgeBox"></asp:TextBox>
    </p>
    <p>
        NI: <asp:TextBox runat="server" ID="NINumberBox"></asp:TextBox>
    </p>
    <p>
        <asp:Button runat="server" ID="SendToServerBtn" Text="Send To Service" OnClick="SendToServerBtn_Click"/>
    </p>
</asp:Content>
