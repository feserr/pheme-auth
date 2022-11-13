import { describe, it, expect } from '@jest/globals';
import request from 'supertest';

const baseUrl = 'http://127.0.0.1:8000';

async function GetCookie() {
  const response = await request(baseUrl)
    .post('/api/auth/login')
    .send({ email: 'test@test.com', password: 'test' });

  return response.get('Set-Cookie')[0];
}

describe('Register endpoint', () => {
  it('register new user missing name', async () => {
    const response = await request(baseUrl)
      .post('/api/auth/register')
      .send({ email: 'test@test.com', password: 'test' });

    expect(response.statusCode).toBe(400);
  });

  it('register new user missing email', async () => {
    const response = await request(baseUrl)
      .post('/api/auth/register')
      .send({ name: 'test', password: 'test' });

    expect(response.statusCode).toBe(400);
  });

  it('register new user missing password', async () => {
    const response = await request(baseUrl)
      .post('/api/auth/register')
      .send({ name: 'test', email: 'test@test.com' });

    expect(response.statusCode).toBe(400);
  });

  it('register new user', async () => {
    const response = await request(baseUrl)
      .post('/api/auth/register')
      .send({ name: 'test', email: 'test@test.com', password: 'test' });

    expect(response.statusCode).toBe(200);
    expect(response.headers['content-type']).toContain('application/json');
    expect(response.body).toHaveProperty('id');
    expect(response.body).toHaveProperty('version');
    expect(response.body).toHaveProperty('userName');
    expect(response.body).toHaveProperty('email');
    expect(response.body).toHaveProperty('createdAt');
  });
});

describe('Register endpoint with data', () => {
  it('register duplicate user', async () => {
    const response = await request(baseUrl)
      .post('/api/auth/register')
      .send({ name: 'test', email: 'test@test.com', password: 'test' });

    expect(response.statusCode).toBe(400);
  });
});

describe('Login endpoint', () => {
  it('login with existing user', async () => {
    const response = await request(baseUrl)
      .post('/api/auth/login')
      .send({ email: 'test@test.com', password: 'test' });

    expect(response.statusCode).toBe(200);
    expect(response.get('Set-Cookie')).toBeDefined();
  });

  it('get logged user info', async () => {
    const cookie = await GetCookie();

    expect(cookie).toBeDefined();

    const response = await request(baseUrl)
      .get('/api/auth/user')
      .set('Cookie', cookie);

    expect(response.statusCode).toBe(200);
  });
});

describe('Logout endpoint', () => {
  it('logout', async () => {
    const cookie = await GetCookie();

    expect(cookie).toBeDefined();

    const response = await request(baseUrl)
      .post('/api/auth/logout')
      .set('Cookie', cookie);

    expect(response.statusCode).toBe(200);
  });
});

describe('Delete endpoint', () => {
  it('Delete un-logged user', async () => {
    const response = await request(baseUrl)
      .delete('/api/auth/user');

    expect(response.statusCode).toBe(401);
  });

  it('Delete logged user', async () => {
    const cookie = await GetCookie();
    expect(cookie).toBeDefined();

    const response = await request(baseUrl)
      .delete('/api/auth/user')
      .set('Cookie', cookie);

    expect(response.statusCode).toBe(200);
  });
});