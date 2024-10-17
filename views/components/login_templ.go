// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"w-full flex justify-center px-4 py-8 xl:px-8\"><div class=\"max-w-5xl mx-auto\"><div class=\"flex flex-col items-center md:flex-row\"><div class=\"w-full space-y-5 md:w-3/5 md:pr-16\"><p class=\"font-medium text-blue-500 uppercase\" data-primary=\"blue-500\">Building Businesses</p><h2 class=\"text-2xl font-extrabold leading-none text-black sm:text-3xl md:text-5xl\">Changing The Way People Do Business.</h2><p class=\"text-xl text-gray-600 md:pr-16\">Learn how to engage with your visitors and teach them about your mission. We're revolutionizing the way customers and businesses interact.</p></div><div class=\"w-full mt-16 md:mt-0 md:w-2/5\"><div class=\"relative z-10 h-auto p-8 py-10 overflow-hidden bg-white border-b-2 border-gray-300 rounded-lg shadow-2xl px-7\" data-rounded=\"rounded-lg\" data-rounded-max=\"rounded-full\"><h3 class=\"mb-6 text-2xl font-medium text-center\">Sign in to your Account</h3><input type=\"text\" name=\"email\" id=\"email\" class=\"block w-full px-4 py-3 mb-4 border border-2 border-transparent border-gray-200 rounded-lg focus:ring focus:ring-blue-500 focus:outline-none\" data-rounded=\"rounded-lg\" data-primary=\"blue-500\" placeholder=\"Email address\"> <input type=\"password\" name=\"password\" id=\"password\" class=\"block w-full px-4 py-3 mb-4 border border-2 border-transparent border-gray-200 rounded-lg focus:ring focus:ring-blue-500 focus:outline-none\" data-rounded=\"rounded-lg\" data-primary=\"blue-500\" placeholder=\"Password\"><div class=\"block\"><button class=\"w-full px-3 py-4 font-medium text-white bg-blue-600 rounded-lg\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\">Log Me In</button></div><p class=\"w-full mt-4 text-sm text-center text-gray-500\">Don't have an account? <a href=\"/signup\" class=\"text-blue-500 underline\">Sign up here</a></p></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate