import {
  Heading,
  Avatar,
  Box,
  Center,
  Stack,
  VStack,
  Button,
  Stat,
  StatLabel,
  StatHelpText,
  Link,
} from '@chakra-ui/react';
import { useCallback, useState } from 'react';
import { ENDPOINT_API_POST_USER } from '../../api/api';
import useLoginState from '../../zustand/todoLogin';
import axios from 'axios';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Toast } from '../../components/02-Reusable/Toast/Toast';

export default function SocialProfileSimple() {
  const navigate = useNavigate();
  const { setIsLoggedOut, setUserId, dataId, setDataId, setMessage, setToken } =
    useLoginState();
  const [user, setUser] = useState([]);

  const HandleLogOut = () => {
    setUserId('');
    setDataId('');
    setMessage('');
    setToken('');
    setIsLoggedOut();
    navigate('/');
    useLoginState.persist.clearStorage();
    localStorage.removeItem('tokenId');
    Toast.fire({
      icon: 'success',
      title: 'Berhasil Keluar',
    });
  };

  const getUser = useCallback(async () => {
    const headers = {
      Authorization: 'Bearer ' + localStorage.getItem('tokenId'),
    };
    const res = await axios.get(ENDPOINT_API_POST_USER + `${dataId}`, {
      headers,
    });
    setUser(res.data);
  }, [dataId]);

  useEffect(() => {
    getUser();
  }, [getUser]);

  return (
    <Center py={6}>
      <Box
        maxW={'320px'}
        w={'full'}
        bg={'gray.75'}
        boxShadow={'2xl'}
        rounded={'lg'}
        p={6}
        textAlign={'center'}
      >
        <Heading lineHeight={1.1} fontSize={{ base: '2xl', sm: '3xl' }}>
          Profile
        </Heading>
        <Avatar
          size={'xl'}
          src={''}
          alt={'Avatar Alt'}
          mb={4}
          pos={'relative'}
          name={`${dataId.nama_depan} ${dataId.nama_belakang}`}
        />
        <VStack justify={'left'} direction={'row'} mt={6}>
          <Stat>
            <StatLabel>{`${
              dataId.nama_depan.charAt(0).toUpperCase() +
              dataId.nama_depan.slice(1)
            } ${
              dataId.nama_belakang.charAt(0).toUpperCase() +
              dataId.nama_belakang.slice(1)
            }`}</StatLabel>
            <StatHelpText>{user.nama_depan}</StatHelpText>
          </Stat>
          <Stat>
            <StatLabel>Pendidikan</StatLabel>
            <StatHelpText>SMA</StatHelpText>
          </Stat>
        </VStack>
        <Center>
          <Stack mt={15} direction={'row'} spacing={4}>
            <Link color="pink.500" href="./editProfile">
              Pengaturan
            </Link>
          </Stack>
        </Center>
        <Stack mt={8} direction={'row'} spacing={4}>
          <Button
            flex={1}
            fontSize={'sm'}
            rounded={'full'}
            bg={'pink.400'}
            color={'white'}
            boxShadow={
              '0px 1px 25px -5px rgb(66 153 225 / 48%), 0 10px 10px -5px rgb(66 153 225 / 43%)'
            }
            _hover={{
              bg: 'pink.400',
            }}
            _focus={{
              bg: 'pink.400',
            }}
          >
            Hapus Akun
          </Button>
          <Button
            flex={1}
            fontSize={'sm'}
            rounded={'full'}
            bg={'pink.400'}
            color={'white'}
            onClick={() => HandleLogOut()}
            boxShadow={
              '0px 1px 25px -5px rgb(66 153 225 / 48%), 0 10px 10px -5px rgb(66 153 225 / 43%)'
            }
            _hover={{
              bg: 'pink.400',
            }}
            _focus={{
              bg: 'pink.400',
            }}
          >
            Keluar
          </Button>
        </Stack>
      </Box>
    </Center>
  );
}
