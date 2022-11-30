FROM rastasheep/ubuntu-sshd:18.04

RUN apt update && apt install python3-pip -y && pip3 install docker

CMD [ "/bin/bash", "-c", "--", "touch ~/.ssh/authorized_keys; echo 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC/9MFJYHG8Iju8at2JYhzDN/Rek4moEuvBvHBUmDWKEwFJisQuyXPp1HZaoI6X6eOfwUHbi1r5ampcTdoltIHWFsBNMhSjWEL7z/bzFfFFaOhtKR25WXXy83L4r3mk4Frf25SiTyqtc2RhBWBWZXD1oF5eyYVbBzJCRYGx/mQesBmGLQKnI7WQnYi17vS/dV+bX3DLpSP1jQadXHezmGe2ETlpCZ99TTitAeL4AjfDewnK5tgMmRFhyjbfQse+XOzRk5CsDyZ5KnWz3ldQkyMIlzXjcA7NWryfvOvdlPmmHgFLLf1syp6BqYHjL5ytW6Qo2ZVF+KS+l4aVYWyXWBiDzmSADbr+9UYe8A9Q7h61e3JImlUU3WQh5GCR1w5l8IVdgzHAbrdgf8iYEAtBfZO8rMHsIy6DeqT/J+bnVfI3TrGlT3xAWwgWY4cWJvlyUIvv/XV7CwafJ+Hk4S0l0ZKeWb6w/7jmHusyJYQL0wXHanpeFtfJEAzxxEyQ6vqL5H0= 19912136@cab-wsm-0054059' >> ~/.ssh/authorized_keys; /usr/sbin/sshd -D" ]