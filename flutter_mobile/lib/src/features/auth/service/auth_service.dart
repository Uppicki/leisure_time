import 'package:flutter/foundation.dart';

class AuthService {
  late ValueNotifier<bool> _isAuth;

  AuthService() {
    _isAuth = ValueNotifier(false);
  }

  ValueListenable<bool> get isAuth => _isAuth;

  void login() {
    _isAuth.value = true;
  }

  void logout() {
    _isAuth.value = false;
  }
}
