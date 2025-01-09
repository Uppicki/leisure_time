import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/screens/login_screen.dart';
import 'package:flutter_mobile/src/screens/reg_screen.dart';

class AuthScreen extends StatelessWidget {
  final Function _callback;

  AuthScreen({
    required Function callback,
  }) : _callback = callback;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 32),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              OutlinedButton(
                onPressed: () async {
                  final res = await Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (_) => LoginScreen(),
                    ),
                  );

                  if (res != null && res) {
                    _callback();
                  }
                },
                child: Text("Login"),
              ),
              OutlinedButton(
                onPressed: () => Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (_) => RegScreen(),
                  ),
                ),
                child: Text("Reg"),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
