package com.wuriyanto.dev.webinar

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Button
import android.widget.ImageView
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.google.android.material.textfield.TextInputEditText
import com.wuriyanto.dev.webinar.model.ProfileResponse
import com.wuriyanto.dev.webinar.model.SignUp
import com.wuriyanto.dev.webinar.model.SignUpResponse
import com.wuriyanto.dev.webinar.service.SignUpService
import com.wuriyanto.dev.webinar.shared.ServiceBuilder
import com.wuriyanto.dev.webinar.shared.removeAccessTokenToPreference
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class SignUpActivity : AppCompatActivity() {

    private val TAG = "SignUpActivity"

    private lateinit var back: ImageView
    private lateinit var textEmail: TextInputEditText
    private lateinit var textFullName: TextInputEditText
    private lateinit var password: TextInputEditText
    private lateinit var buttonSubmit: Button

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_sign_up)

        textEmail = findViewById(R.id.txt_sign_up_email)
        textFullName = findViewById(R.id.txt_sign_up_fullname)
        password = findViewById(R.id.txt_sign_up_password)

        buttonSubmit = findViewById(R.id.btn_submit_signup)

        back = findViewById(R.id.back_from_sign_up)

        back.setOnClickListener {
            startActivity(Intent(this, SignInActivity::class.java))
            finish()
        }

        buttonSubmit.setOnClickListener { submitSignUp(it) }
    }

    fun submitSignUp(view: View) {
        if (textEmail.text.toString().isBlank()) {
            showToast("email cannot be blank")
            return
        }

        if (textFullName.text.toString().isBlank()) {
            showToast("full name cannot be blank")
            return
        }

        if (password.text.toString().isBlank()) {
            showToast("full name cannot be blank")
            return
        }

        val signUpData = SignUp(textEmail.text.toString(), textFullName.text.toString(), password.text.toString())
        val signUpRequest = ServiceBuilder.buildService(SignUpService::class.java)
        signUpRequest.signUp(signUpData).enqueue(
            object : Callback<SignUpResponse> {
                override fun onFailure(call: Call<SignUpResponse>, t: Throwable) {
                    Log.d(TAG, "profile sign up fail: ${t.message}")
                    onResult(null)
                }

                override fun onResponse(call: Call<SignUpResponse>, response: Response<SignUpResponse>) {
                    Log.d(TAG, "profile sign up succeed")
                    onResult(response.body())
                }

            }
        )
    }

    private fun onResult(response: SignUpResponse?) {
        if (response == null) {
            showToast("sign up failed")
        } else {
            if (response.success) {
                removeAccessTokenToPreference(this)
                startActivity(Intent(this, SignInActivity::class.java))
                finish()

            } else {
                showToast("sign up failed")
            }
        }
    }

    private fun showToast(msg: String) {
        Toast.makeText(applicationContext, msg, Toast.LENGTH_SHORT).show()
    }
}