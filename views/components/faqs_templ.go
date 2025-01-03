// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func FAQs() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"relative py-16 bg-white min-w-screen animation-fade animation-delay\"><div class=\"container px-0 px-8 mx-auto sm:px-12 xl:px-5\"><p class=\"text-xs font-bold text-left text-pink-500 uppercase sm:mx-6 sm:text-center sm:text-normal sm:font-bold\" data-primary=\"pink-500\">Got a Question? We’ve got answers.</p><h3 class=\"mt-1 text-2xl font-bold text-left text-gray-800 sm:mx-6 sm:text-3xl md:text-4xl lg:text-5xl sm:text-center sm:mx-0\">Frequently Asked Questions</h3><div class=\"w-full px-6 py-6 mx-auto mt-10 bg-white border border-gray-200 rounded-lg sm:px-8 md:px-12 sm:py-8 sm:shadow lg:w-5/6 xl:w-2/3\" data-rounded=\"rounded-lg\" data-rounded-max=\"rounded-full\"><h3 class=\"text-lg font-bold text-purple-500 sm:text-xl md:text-2xl\" data-primary=\"purple-500\">How does it work?</h3><p class=\"mt-2 text-base text-gray-600 sm:text-lg md:text-normal\">Our platform works with your content to provides insights and metrics on how you can grow your business and scale your infastructure.</p></div><div class=\"w-full px-6 py-6 mx-auto mt-10 bg-white border border-gray-200 rounded-lg sm:px-8 md:px-12 sm:py-8 sm:shadow lg:w-5/6 xl:w-2/3\" data-rounded=\"rounded-lg\" data-rounded-max=\"rounded-full\"><h3 class=\"text-lg font-bold text-purple-500 sm:text-xl md:text-2xl\" data-primary=\"purple-500\">Do you offer team pricing?</h3><p class=\"mt-2 text-base text-gray-600 sm:text-lg md:text-normal\">Yes, we do! Team pricing is available for any plan. You can take advantage of 30% off for signing up for team pricing of 10 users or more.</p></div><div class=\"w-full px-6 py-6 mx-auto mt-10 bg-white border border-gray-200 rounded-lg sm:px-8 md:px-12 sm:py-8 sm:shadow lg:w-5/6 xl:w-2/3\" data-rounded=\"rounded-lg\" data-rounded-max=\"rounded-full\"><h3 class=\"text-lg font-bold text-purple-500 sm:text-xl md:text-2xl\" data-primary=\"purple-500\">How do I make changes and configure my site?</h3><p class=\"mt-2 text-base text-gray-600 sm:text-lg md:text-normal\">You can easily change your site settings inside of your site dashboard by clicking the top right menu and clicking the settings button.</p></div><div class=\"w-full px-6 py-6 mx-auto mt-10 bg-white border border-gray-200 rounded-lg sm:px-8 md:px-12 sm:py-8 sm:shadow lg:w-5/6 xl:w-2/3\" data-rounded=\"rounded-lg\" data-rounded-max=\"rounded-full\"><h3 class=\"text-lg font-bold text-purple-500 sm:text-xl md:text-2xl\" data-primary=\"purple-500\">How do I add a custom domain?</h3><p class=\"mt-2 text-base text-gray-600 sm:text-lg md:text-normal\">You can easily change your site settings inside of your site dashboard by clicking the top right menu and clicking the settings button.</p></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
