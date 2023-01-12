import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';
import {
  Box,
  Button,
  Checkbox,
  Container,
  FormControl,
  FormLabel,
  Heading,
  HStack,
  IconButton,
  Image,
  Input,
  InputGroup,
  InputRightElement,
  Stack,
  Text,
  useBreakpointValue,
  useColorModeValue,
} from '@chakra-ui/react';
import { Link, useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { Toast } from '../../components/02-Reusable/Toast/Toast';
import axios from 'axios';
import useLoginState from '../../zustand/todoLogin';

export default function Login(props) {
  const navigate = useNavigate();
  const { isLoggedIn, setIsLoggedIn, setUserId } = useLoginState();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [passwordType, setPasswordType] = useState(false);

  const HandleSubmit = async (e) => {
    e.preventDefault();

    const user = {
      email,
      password,
    };

    await axios
      .post('http://localhost:8080/login', user, {
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      })
      .then((response) => {
        setIsLoggedIn(true);
        Toast.fire({
          icon: 'success',
          title: 'Berhasil Masuk',
        });
        const UserID = response?.data?.user;
        setUserId(UserID);
        const accessToken = response?.data?.token;
        localStorage.setItem('token', accessToken);
      })
      .catch((error) => {
        Toast.fire({
          icon: 'error',
          title: 'Gagal Masuk, Email atau Password salah',
        });
      });
  };

  useEffect(() => {
    if (isLoggedIn) {
      navigate('/');
    }
  }, [isLoggedIn, navigate]);

  return (
    <Container>
      <form onSubmit={HandleSubmit}>
        <Container
          data-aos="flip-up"
          maxW="lg"
          py={{
            base: '12',
            md: '8',
          }}
          px={{
            base: '0',
            sm: '8',
          }}
        >
          <Stack
            spacing={8}
            backgroundColor={useColorModeValue('#F5FFFA', '#696969')}
          >
            <Stack spacing={6}>
              <Image
                src={`${process.env.PUBLIC_URL + `/image/logo.svg`}`}
                alt="logo"
                style={{
                  height: '20',
                  width: '20',
                  display: 'block',
                  margin: 'auto',
                }}
              />
              <Stack
                spacing={{
                  base: '2',
                  md: '3',
                }}
                textAlign="center"
              >
                <Heading
                  size={useBreakpointValue({
                    base: 'xs',
                    md: 'sm',
                  })}
                >
                  Masuk ke Akun Anda
                </Heading>
              </Stack>
            </Stack>
            <Box
              py={{
                base: '0',
                sm: '8',
              }}
              px={{
                base: '4',
                sm: '10',
              }}
              bg={useBreakpointValue({
                base: 'transparent',
                sm: 'bg-surface',
              })}
              boxShadow={{
                base: 'none',
                sm: useColorModeValue('md', 'md-dark'),
              }}
              borderRadius={{
                base: 'none',
                sm: 'xl',
              }}
            >
              <Stack spacing={6}>
                <Stack spacing={5}>
                  <FormControl isRequired>
                    <FormLabel htmlFor="email">Email</FormLabel>
                    <Input
                      id="email"
                      type="email"
                      aria-label="email"
                      onChange={(e) => setEmail(e.target.value)}
                      value={email}
                      placeholder="Masukan Email Anda"
                      focusBorderColor={useColorModeValue(
                        'accentLight.400',
                        'accentDark.400'
                      )}
                      required
                    />
                  </FormControl>
                  <FormControl isRequired>
                    <FormLabel>Kata Sandi</FormLabel>
                    <InputGroup size={'lg'}>
                      <Input
                        type={passwordType ? 'text' : 'password'}
                        id="password"
                        aria-label="password"
                        onChange={(e) => setPassword(e.target.value)}
                        value={password}
                        autoComplete="current-password"
                        focusBorderColor={useColorModeValue(
                          'accentLight.400',
                          'accentDark.400'
                        )}
                        required
                      />
                      <InputRightElement width="4.5rem">
                        <IconButton
                          icon={passwordType ? <ViewIcon /> : <ViewOffIcon />}
                          onClick={() => setPasswordType(!passwordType)}
                          variant="ghost"
                        />
                      </InputRightElement>
                    </InputGroup>
                  </FormControl>
                  <HStack justify={'space-between'}>
                    <Checkbox>Ingat Saya</Checkbox>
                    <Button
                      variant={'link'}
                      fontSize={'sm'}
                      colorScheme={'blue'}
                      as={Link}
                      to="/lupa-password"
                    >
                      Lupa Kata Sandi?
                    </Button>
                  </HStack>
                </Stack>
                <Stack spacing={6}>
                  <Button
                    type="submit"
                    color={useColorModeValue('white', 'black')}
                    bg={useColorModeValue('accentLight.400', 'accentDark.400')}
                    _hover={{
                      bg: useColorModeValue(
                        'accentLight.500',
                        'accentDark.500'
                      ),
                    }}
                  >
                    Masuk
                  </Button>
                  <Text textAlign={'center'}>
                    Belum memiliki akun?{' '}
                    <Link
                      color={useColorModeValue(
                        'accentLight.400',
                        'accentDark.400'
                      )}
                      to="/mendaftar"
                    >
                      Daftar
                    </Link>
                  </Text>
                </Stack>
              </Stack>
            </Box>
          </Stack>
        </Container>
      </form>
    </Container>
  );
}
